package main

import (
	"api/model"
	"api/query"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var SendToQueue = true

func main() {
	gormDB, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&model.Customer{}, &model.Account{}, &model.Tx{})
	query.SetDefault(gormDB)

	app := fiber.New()

	store := session.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	queue := "transaction"

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)

	gQInput := &sqs.GetQueueUrlInput{
		QueueName: &queue,
	}

	result, err := GetQueueURL(context.TODO(), client, gQInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	app.Post("/register", func(c *fiber.Ctx) error {
		body := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}
		pHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			return err
		}
		if err := query.Customer.Create(&model.Customer{
			Name:     body.Name,
			Email:    body.Email,
			Password: string(pHash),
			Accounts: nil,
		}); err != nil {
			return err
		}
		return c.SendStatus(http.StatusOK)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		body := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}
		e := fiber.NewError(http.StatusBadRequest, "password wrong")
		cus, err := query.Customer.Where(query.Customer.Email.Eq(body.Email)).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return e
			}
			return err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(cus.Password), []byte(body.Password)); err != nil {
			return e
		}

		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		sess.Set("cid", cus.ID)

		// Save session
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return c.SendStatus(http.StatusOK)
	})

	app.Post("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		// Save session
		if err := sess.Destroy(); err != nil {
			panic(err)
		}
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		cus, err := query.Customer.Where(query.Customer.ID.Eq(cid)).First()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(cus.JSON())
	})

	app.Get("/accounts", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		acc, err := query.Account.Where(query.Account.CustomerID.Eq(cid)).Find()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(x(acc).JSON())
	})

	app.Put("/account", func(c *fiber.Ctx) error {
		body := struct {
			Name string `json:"name"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		if err := query.Account.Create(&model.Account{
			CustomerID: cid,
			Name:       body.Name,
			Balance:    0,
		}); err != nil {
			return err
		}

		return c.SendStatus(http.StatusOK)
	})

	app.Post("/pre-transfer", func(c *fiber.Ctx) error {
		body := struct {
			SourceAccId uint    `json:"source_acc_id"`
			AccNo       string  `json:"acc_no"`
			Amount      float64 `json:"amount"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		sourceAcc, err := query.Account.Where(query.Account.ID.Eq(body.SourceAccId), query.Account.CustomerID.Eq(cid)).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "source acc not found")
			}
			return err
		}

		if sourceAcc.Balance < body.Amount {
			return fiber.NewError(http.StatusBadRequest, "balance not enough")
		}

		targetAcc, err := query.Account.Where(query.Account.No.Eq(body.AccNo)).Preload(query.Account.Customer).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "target acc not found")
			}
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"target_name": targetAcc.Customer.Name,
		})
	})

	app.Post("/transfer", func(c *fiber.Ctx) error {
		body := struct {
			SourceAccId uint    `json:"source_acc_id"`
			AccNo       string  `json:"acc_no"`
			Amount      float64 `json:"amount"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		sourceAcc, err := query.Account.Where(query.Account.ID.Eq(body.SourceAccId), query.Account.CustomerID.Eq(cid)).First()
		if err != nil {
			return err
		}

		if sourceAcc.Balance < body.Amount {
			return fiber.NewError(http.StatusBadRequest, "balance not enough")
		}

		targetAcc, err := query.Account.Where(query.Account.No.Eq(body.AccNo)).Preload(query.Account.Customer).First()
		if err != nil {
			return err
		}

		t := model.Tx{
			AccountID:    sourceAcc.ID,
			DesAccountID: targetAcc.ID,
			Amount:       body.Amount * 1,
			//Title:        "Transfer",
			//Description:  fmt.Sprintf("Transfer to account %s", masker.String(masker.MID, targetAcc.No)),
			Status: model.TransactionPending,
		}

		if !SendToQueue {
			if err := query.Q.Transaction(func(tx *query.Query) error {
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(targetAcc.ID)).UpdateSimple(tx.Account.Balance.Add(body.Amount)); err != nil {
					return err
				}
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(sourceAcc.ID), tx.Account.Balance.Gte(body.Amount)).UpdateSimple(tx.Account.Balance.Sub(body.Amount)); err != nil {
					return err
				}
				t.Status = model.TransactionSuccess
				if err := tx.WithContext(context.Background()).Tx.Create(&t); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
		} else {
			if err := query.Tx.Create(&t); err != nil {
				return err
			}

			sMInput := &sqs.SendMessageInput{
				QueueUrl:     queueURL,
				DelaySeconds: 1,
				MessageAttributes: map[string]types.MessageAttributeValue{
					"TxID": {
						DataType:    aws.String("String"),
						StringValue: aws.String(fmt.Sprintf("%d", t.ID)),
					},
				},
				MessageBody: aws.String("HellO!"),
			}

			resp, err := SendMsg(context.TODO(), client, sMInput)
			if err != nil {
				fmt.Println("Got an error sending the message:")
				fmt.Println(err)
				return err
			}

			fmt.Println("Sent message with ID: " + *resp.MessageId)
		}

		return c.SendStatus(http.StatusOK)
	})

	app.Listen(":3003")
}

func x(a model.AccountList) model.AccountList {
	return a
}

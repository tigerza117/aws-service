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
	esql "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var SendToQueue = false

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var baseUrl = os.Getenv("BASE_URL")

	gormDB, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")))
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&model.Customer{}, &model.Account{}, &model.Tx{})
	query.SetDefault(gormDB)

	app := fiber.New()
	app.Use(logger.New())
	var storage fiber.Storage
	if os.Getenv("REDIS_ENABLE") == "TRUE" {
		storage = redis.New(redis.Config{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     6379,
			Database: 0,
			Reset:    false,
			PoolSize: 10 * runtime.GOMAXPROCS(0),
		})
	}

	storeCfg := session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.UUID,
		CookieSameSite: "None",
		Storage:        storage,
	}
	if os.Getenv("RUN_MODE") == "TEST" {
		storeCfg.CookieSameSite = "Lax"
	}
	store := session.New(storeCfg)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	api := app.Group(baseUrl)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is GET -baseUrl-")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is GET /")
	})

	if os.Getenv("SQS_ENABLE") == "TRUE" {
		SendToQueue = true
	}

	var queueURL *string
	var client *sqs.Client

	if SendToQueue {
		queue := os.Getenv("SQS_NAME")

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(os.Getenv("SQS_REGION")),
		)
		if err != nil {
			panic("configuration error, " + err.Error())
		}

		client = sqs.NewFromConfig(cfg)

		gQInput := &sqs.GetQueueUrlInput{
			QueueName: &queue,
		}

		result, err := GetQueueURL(context.TODO(), client, gQInput)
		if err != nil {
			fmt.Println("Got an error getting the queue URL:")
			fmt.Println(err)
			return
		}

		queueURL = result.QueueUrl
	}

	api.Post("/register", func(c *fiber.Ctx) error {
		body := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Check if some parameters are empty
		if body.Email == "" {
			return fiber.NewError(http.StatusBadRequest, "E-mail cannot be empty")
		}
		if body.Name == "" {
			return fiber.NewError(http.StatusBadRequest, "Name cannot be empty")
		}
		if body.Password == "" {
			return fiber.NewError(http.StatusBadRequest, "Password cannot be empty")
		}

		pHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		if err := query.Customer.Create(&model.Customer{
			Name:     body.Name,
			Email:    body.Email,
			Password: string(pHash),
			Accounts: nil,
		}); err != nil {
			var mysqlErr *esql.MySQLError
			if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
				return fiber.NewError(http.StatusBadRequest, "E-mail already used")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return c.SendStatus(http.StatusOK)
	})

	api.Post("/login", func(c *fiber.Ctx) error {
		body := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Check if some parameters are empty
		if body.Email == "" {
			return fiber.NewError(http.StatusBadRequest, "E-mail cannot left blank")
		}
		if body.Password == "" {
			return fiber.NewError(http.StatusBadRequest, "Password cannot left blank")
		}

		e := fiber.NewError(http.StatusBadRequest, "Incorrect E-mail/Password")
		cus, err := query.Customer.Where(query.Customer.Email.Eq(body.Email)).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return e
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(cus.Password), []byte(body.Password)); err != nil {
			return e
		}

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		sess.Set("cid", cus.ID)

		// Save session
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return c.SendStatus(http.StatusOK)
	})

	api.Post("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Save session
		if err := sess.Destroy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return c.SendStatus(http.StatusOK)
	})

	api.Get("/profile", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		cus, err := query.Customer.Where(query.Customer.ID.Eq(cid)).First()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.Status(http.StatusOK).JSON(cus.JSON())
	})

	api.Get("/accounts", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		acc, err := query.Account.Where(query.Account.CustomerID.Eq(cid)).Find()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.Status(http.StatusOK).JSON(x(acc).JSON())
	})

	api.Put("/account", func(c *fiber.Ctx) error {
		body := struct {
			Name string `json:"name"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Check if parameter is empty
		if body.Name == "" {
			return fiber.NewError(http.StatusBadRequest, "Account Name cannot be empty")
		}

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
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
			var mysqlErr *esql.MySQLError
			if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
				return fiber.NewError(http.StatusBadRequest, "Account Name is already used")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.SendStatus(http.StatusOK)
	})

	api.Post("/pre-transfer", func(c *fiber.Ctx) error {
		body := struct {
			Id     uint    `json:"id"`  // Src Account ID
			Acc    string  `json:"acc"` // Dst Account No
			Amount float64 `json:"amount"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		// Check if parameter is empty
		if body.Acc == "" {
			return fiber.NewError(http.StatusBadRequest, "Source Account cannot left blank")
		}
		if body.Amount <= 0 {
			return fiber.NewError(http.StatusBadRequest, "amount cant be negative/zero")
		}

		sourceAcc, err := query.Account.Where(query.Account.ID.Eq(body.Id), query.Account.CustomerID.Eq(cid)).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "source acc not found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		if sourceAcc.Balance < body.Amount {
			return fiber.NewError(http.StatusBadRequest, "balance not enough")
		}

		targetAcc, err := query.Account.Where(query.Account.No.Eq(body.Acc)).Preload(query.Account.Customer).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "target acc not found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"target_account": targetAcc.Name,
			"target_name":    targetAcc.Customer.Name,
		})
	})

	api.Post("/transfer", func(c *fiber.Ctx) error {
		body := struct {
			Id     uint    `json:"id"`  // Src Account ID
			Acc    string  `json:"acc"` // Dst Account No
			Amount float64 `json:"amount"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		cid, ok := sess.Get("cid").(uint)
		if !ok {
			return c.SendStatus(http.StatusForbidden)
		}

		// Check if parameter is empty
		if body.Acc == "" {
			return fiber.NewError(http.StatusBadRequest, "Source Account cannot left blank")
		}
		if body.Amount <= 0 {
			return fiber.NewError(http.StatusBadRequest, "amount cant be negative/zero")
		}

		sourceAcc, err := query.Account.Where(query.Account.ID.Eq(body.Id), query.Account.CustomerID.Eq(cid)).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "source acc not found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		if sourceAcc.Balance < body.Amount {
			return fiber.NewError(http.StatusBadRequest, "balance not enough")
		}

		targetAcc, err := query.Account.Where(query.Account.No.Eq(body.Acc)).Preload(query.Account.Customer).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "target acc not found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		t := model.Tx{
			AccountID:    &sourceAcc.ID,
			DstAccountID: targetAcc.ID,
			Amount:       body.Amount * 1,
			//Title:        "Transfer",
			//Description:  fmt.Sprintf("Transfer to account %s", masker.String(masker.MID, targetAcc.No)),
			Status: model.TransactionPending,
		}

		if !SendToQueue {
			if err := query.Q.Transaction(func(tx *query.Query) error {
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(targetAcc.ID)).UpdateSimple(tx.Account.Balance.Add(body.Amount)); err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(sourceAcc.ID), tx.Account.Balance.Gte(body.Amount)).UpdateSimple(tx.Account.Balance.Sub(body.Amount)); err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}
				t.Status = model.TransactionSuccess
				if err := tx.WithContext(context.Background()).Tx.Create(&t); err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}
				return nil
			}); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		} else {
			if err := query.Tx.Create(&t); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			sMInput := &sqs.SendMessageInput{
				QueueUrl: queueURL,
				MessageAttributes: map[string]types.MessageAttributeValue{
					"TxID": {
						DataType:    aws.String("String"),
						StringValue: aws.String(fmt.Sprintf("%d", t.ID)),
					},
				},
				MessageBody:    aws.String(fmt.Sprintf("Transaction ID: %d", t.ID)),
				MessageGroupId: aws.String("Group1"),
			}

			resp, err := SendMsg(context.TODO(), client, sMInput)
			if err != nil {
				fmt.Println("Got an error sending the message:")
				fmt.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			fmt.Println("Sent message with ID: " + *resp.MessageId)
		}

		return c.SendStatus(http.StatusOK)
	})

	api.Post("/pre-deposit", func(c *fiber.Ctx) error {
		body := struct {
			Acc    string  `json:"acc"` // AccountNo
			Amount float64 `json:"amount"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Check if parameter is empty
		if body.Acc == "" {
			return fiber.NewError(http.StatusBadRequest, "Source Account cannot left blank")
		}
		if body.Amount <= 0 {
			return fiber.NewError(http.StatusBadRequest, "amount cant be negative/zero")
		}

		targetAcc, err := query.Account.Where(query.Account.No.Eq(body.Acc)).Preload(query.Account.Customer).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "target acc not found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"target_account": targetAcc.Name,
			"target_name":    targetAcc.Customer.Name,
		})
	})

	api.Post("/deposit", func(c *fiber.Ctx) error {
		body := struct {
			Acc    string  `json:"acc"` // AccountNo
			Amount float64 `json:"amount"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Check if parameter is empty
		if body.Acc == "" {
			return fiber.NewError(http.StatusBadRequest, "Source Account cannot left blank")
		}
		if body.Amount <= 0 {
			return fiber.NewError(http.StatusBadRequest, "amount cant be negative/zero")
		}

		targetAcc, err := query.Account.Where(query.Account.No.Eq(body.Acc)).Preload(query.Account.Customer).First()
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				return fiber.NewError(http.StatusBadRequest, "target acc not found")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		t := model.Tx{
			DstAccountID: targetAcc.ID,
			Amount:       body.Amount * 1,
			Status:       model.TransactionPending,
		}

		if !SendToQueue {
			if err := query.Q.Transaction(func(tx *query.Query) error {
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(targetAcc.ID)).UpdateSimple(tx.Account.Balance.Add(body.Amount)); err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}
				t.Status = model.TransactionSuccess
				if err := tx.WithContext(context.Background()).Tx.Create(&t); err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}
				return nil
			}); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		} else {
			if err := query.Tx.Create(&t); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			sMInput := &sqs.SendMessageInput{
				QueueUrl: queueURL,
				MessageAttributes: map[string]types.MessageAttributeValue{
					"TxID": {
						DataType:    aws.String("String"),
						StringValue: aws.String(fmt.Sprintf("%d", t.ID)),
					},
				},
				MessageBody:    aws.String(fmt.Sprintf("Transaction ID: %d", t.ID)),
				MessageGroupId: aws.String("Group2"),
			}

			resp, err := SendMsg(context.TODO(), client, sMInput)
			if err != nil {
				fmt.Println("Got an error sending the message:")
				fmt.Println(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			fmt.Println("Sent message with ID: " + *resp.MessageId)
		}

		return c.SendStatus(http.StatusOK)
	})

	app.Listen(os.Getenv("LISTEN"))
}

func x(a model.AccountList) model.AccountList {
	return a
}

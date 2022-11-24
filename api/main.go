package main

import (
	"api/model"
	"api/query"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

const SEND_TO_QUEUE = false

func main() {
	gormdb, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	gormdb.AutoMigrate(&model.Customer{}, &model.Account{})
	query.SetDefault(gormdb)

	app := fiber.New()

	store := session.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

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

		if !SEND_TO_QUEUE {
			if err := query.Q.Transaction(func(tx *query.Query) error {
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(targetAcc.ID)).UpdateSimple(tx.Account.Balance.Add(body.Amount)); err != nil {
					return err
				}
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(sourceAcc.ID), tx.Account.Balance.Gte(body.Amount)).UpdateSimple(tx.Account.Balance.Sub(body.Amount)); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
		}

		return c.SendStatus(http.StatusOK)
	})

	app.Listen(":3000")
}

func x(a model.AccountList) model.AccountList {
	return a
}

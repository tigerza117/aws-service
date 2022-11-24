package main

import (
	"api/model"
	"api/query"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

const JWT_KEY = ""

func main() {
	gormdb, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	gormdb.AutoMigrate(&model.Customer{}, &model.Account{})
	query.SetDefault(gormdb)

	app := fiber.New()

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
		return c.SendStatus(http.StatusOK)
	})

	app.Put("/account", func(c *fiber.Ctx) error {
		body := struct {
			Name string `json:"name"`
		}{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		if err := query.Account.Create(&model.Account{
			CustomerID: 0,
			Name:       body.Name,
			Balance:    0,
		}); err != nil {
			return err
		}

		return c.SendStatus(http.StatusOK)
	})

	app.Listen(":3000")
}

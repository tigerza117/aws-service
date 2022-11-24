package model

import (
	"api/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Customer struct {
	gorm.Model

	Name     string      `gorm:"type:varchar(100);not null"`
	Email    string      `json:"-" gorm:"type:varchar(100);uniqueIndex"`
	Password string      `json:"-"`
	Accounts AccountList `json:"-"`
}

func (c *Customer) JSON() fiber.Map {
	return fiber.Map{
		"id":         c.ID,
		"created_at": c.CreatedAt,
		"updated_at": c.UpdatedAt,
		"name":       c.Name,
		"email":      c.Email,
		"account":    c.Accounts.JSON(),
	}
}

type AccountList []*Account

func (a AccountList) JSON() []fiber.Map {
	return lo.Map[*Account, fiber.Map](a, func(x *Account, y int) fiber.Map {
		return x.JSON()
	})
}

type Account struct {
	gorm.Model

	CustomerID uint    `json:"-"`
	Name       string  `json:"-"`
	No         string  `json:"-"`
	Balance    float64 `json:"-"`

	Customer *Customer `json:"-"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	if a.No == "" {
		rand.Seed(time.Now().UnixNano())
		a.No = helper.RandStringRunes(10)
	}
	return
}

func (a *Account) JSON() fiber.Map {
	return fiber.Map{
		"id":         a.ID,
		"created_at": a.CreatedAt,
		"updated_at": a.UpdatedAt,
		"name":       a.Name,
		"no":         a.No,
		"balance":    a.Balance,
	}
}

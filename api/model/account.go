package model

import (
	"api/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Account struct {
	gorm.Model

	CustomerID uint      `json:"-"`
	Name       string    `json:"-" gorm:"type:varchar(100);uniqueIndex"`
	No         string    `json:"-"`
	Balance    float64   `json:"-"`
	Customer   *Customer `json:"-"`
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

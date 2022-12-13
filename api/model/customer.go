package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
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

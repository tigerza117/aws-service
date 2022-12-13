package model

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionStatus int

const (
	TransactionPending = iota
	TransactionSuccess
	TransactionFail
)

type Tx struct {
	gorm.Model

	AccountID    uint              `json:"-"`
	DstAccountID uint              `json:"-"`
	Amount       float64           `json:"amount"`
	Status       TransactionStatus `json:"status"`

	Account    *Account `json:"-"`
	DstAccount *Account `json:"-"`
}

func (t *Tx) JSON() fiber.Map {
	json := fiber.Map{
		"id":     t.ID,
		"amount": t.Amount,
		"status": t.Status,
	}
	return json
}

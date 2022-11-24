package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model

	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Password string    `json:"-"`
	Accounts []Account `json:"-"`
}

type Account struct {
	gorm.Model

	CustomerID uint    `json:"customer_id"`
	Name       string  `json:"name"`
	Balance    float64 `json:"balance"`
}

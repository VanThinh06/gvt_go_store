// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Type string

const (
	TypeAdmin Type = "admin"
	TypeUser  Type = "user"
)

type Category struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	National  string         `json:"national"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"update_at"`
}

type Order struct {
	ID            uuid.UUID      `json:"id"`
	IDTransaction uuid.UUID      `json:"id_transaction"`
	IDProduct     uuid.UUID      `json:"id_product"`
	Amount         null.Int  `json:"amount"`
	Data           null.String `json:"data"`
	Status         null.Int  `json:"status"`
	Qty            null.Int  `json:"qty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdateAt      time.Time      `json:"update_at"`
}

type Product struct {
	ID          uuid.UUID      `json:"id"`
	IDCategory  uuid.UUID      `json:"id_category"`
	Name        string         `json:"name"`
	Price        null.Int  `json:"price"`
	Image        null.String `json:"image"`
	ListImage   []string       `json:"list_image"`
	Description  null.String `json:"description"`
	Sold         null.Int  `json:"sold"`
	Status       null.Int  `json:"status"`
	Sale         null.Int  `json:"sale"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdateAt    time.Time      `json:"update_at"`
}

type Transaction struct {
	ID        uuid.UUID      `json:"id"`
	IDUser    uuid.UUID      `json:"id_user"`
	Status     null.Int  `json:"status"`
	Amount     null.Int  `json:"amount"`
	Message    null.String `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"update_at"`
}

type User struct {
	ID            uuid.UUID      `json:"id"`
	Name           null.String `json:"name"`
	Address        null.String `json:"address"`
	Phone          null.String `json:"phone"`
	Email         string         `json:"email"`
	TypeUser      Type           `json:"type_user"`
	Password      string         `json:"password"`
	Payment        null.String `json:"payment"`
	PaymentInfo    null.String `json:"payment_info"`
	PaymentNumber  null.String `json:"payment_number"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdateAt      time.Time      `json:"update_at"`
}
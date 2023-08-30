package models

import (
	"time"
)

type Customer struct {
	Customer_Id int32     `json:"customer_id" bson:"customer_id"`
	FirstName   string    `json:"first_name" bson:"first_name"`
	LastName    string    `json:"last_name" bson:"last_name"`
	BankId      string    `json:"bank_id" bson:"bank_id"`
	Balance     float32   `json:"balance" bson:"balance"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	IsActive    bool      `json:"is_active" bson:"is_active"`
}

type CustomerResponse struct {
	CustomerId int32     `json:"customer_id" bson:"customer_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

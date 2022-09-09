package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	User_id      int    `json:"user_id"`
	Payment_slip string `json:"payment_slip"`
	Total_price  int    `json:"total_price"`
	Status       string `json:"status"`
}

type Transaction_Detail struct {
	gorm.Model
	Transaction_id int    `json:"transaction_id"`
	Product_name   string `json:"product_name"`
	Quantity       int    `json:"quantity"`
	Price          int    `json:"price"`
}

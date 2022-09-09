package entities

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User_id    int `json:"user_id"`
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type CartDetail struct {
	User_id       int
	Product_id    int
	Quantity      int
	Product_name  string
	Product_price int
	Total_price   int
}

type CartProduct struct {
	Name  string
	Price int
}

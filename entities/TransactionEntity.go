package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	User_id      int
	Payment_slip string
	Total_price  int
	Status       string
}

package entities

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User_id    int
	Product_id int
	Quantity   int
}

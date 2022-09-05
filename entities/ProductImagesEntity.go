package entities

import (
	"gorm.io/gorm"
)

type ProductImages struct {
	gorm.Model
	Image      string
	Product_id int
}

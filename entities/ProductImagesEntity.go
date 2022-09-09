package entities

import (
	"gorm.io/gorm"
)

type ProductImages struct {
	gorm.Model
	Image      string `json:"image"`
	Product_id int    `json:"product_id"`
}

type ProductImagesLite struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
}

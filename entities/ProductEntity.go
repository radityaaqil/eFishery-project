package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
}

type ProductDetail struct {
	Name        string              `json:"name"`
	Category    string              `json:"category"`
	Price       int                 `json:"price"`
	Images      []ProductImagesLite `json:"images"`
	Unit        string              `json:"unit"`
	Description ProductDescription  `json:"description"`
}

type ProductList struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Image    string `json:"image"`
}

type ProductDescription struct {
	Principal string
	Usage     string
}

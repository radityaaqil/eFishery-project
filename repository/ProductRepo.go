package repository

import (
	"efishery-project/entities"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	GetListProduct(page int, category string, minPrice int, maxPrice int) ([]entities.ProductList, error)
	GetProductDetail(id int) (entities.ProductDetail, error)
}

type Repository_Product struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *Repository_Product {
	return &Repository_Product{db}
}

func (r *Repository_Product) GetListProduct(page int, category string, minPrice int, maxPrice int) ([]entities.ProductList, error) {
	var products []entities.Product
	var productImages []entities.ProductImages
	var productList []entities.ProductList

	limit := 5
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	paginate := fmt.Sprintf("limit %d offset %d", limit, offset)

	var catQuery string
	if category != "" {
		catQuery = fmt.Sprintf("AND category LIKE '%s'", category)
	}

	var priceRange string
	if minPrice <= 0 && maxPrice <= 0 {
		priceRange = ""
	} else if minPrice > 0 && maxPrice <= 0 {
		priceRange = fmt.Sprintf("AND price >= %d", minPrice)
	} else if minPrice <= 0 && maxPrice > 0 {
		priceRange = fmt.Sprintf("AND price <= %d", maxPrice)
	} else if minPrice > 0 && maxPrice > 0 {
		priceRange = fmt.Sprintf("AND price BETWEEN %d and %d", minPrice, maxPrice)
	}

	query := fmt.Sprintf("select * from products where true %s %s %s", catQuery, priceRange, paginate)
	// resultProduct := r.db.Where("category = ?", category).Limit(limit).Offset(offset).Find(&products)
	// resultProduct := r.db.Raw("select * from products where category = ? and price between ? and ? limit ? offset ?", category, minPrice, maxPrice, limit, offset).Scan(&products)
	resultProduct := r.db.Raw(query).Scan(&products)

	errProduct := resultProduct.Error

	if errProduct != nil {
		return productList, errProduct
	}

	log.Println(products)

	for i := 0; i < len(products); i++ {
		var productImage entities.ProductImages
		resultImage := r.db.First(&productImage, products[i].ID)
		errResultImage := resultImage.Error

		if errResultImage != nil {
			return productList, errResultImage
		}

		productImages = append(productImages, productImage)
	}

	for i := 0; i < len(products); i++ {
		prod := entities.ProductList{products[i].Name, products[i].Category, products[i].Price, productImages[i].Image}

		productList = append(productList, prod)
	}

	return productList, nil
}

func (r *Repository_Product) GetProductDetail(id int) (entities.ProductDetail, error) {
	var product entities.Product
	var productDetail entities.ProductDetail
	var productImages []entities.ProductImages
	var productImagesLite []entities.ProductImagesLite

	resultProduct := r.db.First(&product, id)

	errResultProduct := resultProduct.Error

	if errResultProduct != nil {
		return productDetail, errResultProduct
	}

	resultImages := r.db.Where("product_id = ?", product.ID).Find(&productImages)

	errResultImages := resultImages.Error

	if errResultImages != nil {
		return productDetail, errResultImages
	}

	for i := 0; i < len(productImages); i++ {
		value := entities.ProductImagesLite{int(productImages[i].ID), productImages[i].Image}
		productImagesLite = append(productImagesLite, value)
	}

	var description entities.ProductDescription
	json.Unmarshal([]byte(product.Description), &description)

	productDetail = entities.ProductDetail{product.Name, product.Category, product.Price, productImagesLite, product.Unit, description}

	return productDetail, nil

}

package repository

import (
	"efishery-project/entities"
	"errors"
	"log"

	"gorm.io/gorm"
)

type RepositoryCart interface {
	InsertToCart(cart entities.Cart) (entities.Cart, error)
	GetCart(id int) ([]entities.Cart, error)
}

type Repository_Cart struct {
	db *gorm.DB
}

func NewRepositoryCart(db *gorm.DB) *Repository_Cart {
	return &Repository_Cart{db}
}

func (r *Repository_Cart) InsertToCart(cart entities.Cart) (entities.Cart, error) {
	var haveProduct entities.Cart
	var insertCart entities.Cart
	var findProduct entities.Product

	resultProduct := r.db.First(&findProduct, cart.Product_id)

	resultproductRowsAffected := resultProduct.RowsAffected

	errResProd := "Such product unavailable"

	if resultproductRowsAffected == 0 {
		return haveProduct, errors.New(errResProd)
	}

	resultHaveProd := r.db.Where("user_id = ? AND product_id = ?", cart.User_id, cart.Product_id).Find(&haveProduct)

	rowAffected := resultHaveProd.RowsAffected

	if rowAffected != 0 {
		newQuantity := haveProduct.Quantity + cart.Quantity
		updateQuantity := r.db.Model(&entities.Cart{}).Where("user_id = ? and product_id = ?", haveProduct.User_id, haveProduct.Product_id).Update("quantity", newQuantity)

		errUpdateQuantity := updateQuantity.Error

		if errUpdateQuantity != nil {
			return haveProduct, errUpdateQuantity
		}

		var updatedCart entities.Cart
		getUpdatedCart := r.db.Where("user_id = ? AND product_id = ?", cart.User_id, cart.Product_id).Find(&updatedCart)

		errGetUpdatedCart := getUpdatedCart.Error

		if errGetUpdatedCart != nil {
			return updatedCart, errGetUpdatedCart
		}

		return updatedCart, nil
	} else {
		insertCart = entities.Cart{cart.Model, cart.User_id, cart.Product_id, cart.Quantity}
		insert := r.db.Create(&insertCart)

		errInsert := insert.Error

		if errInsert != nil {
			log.Panic("Error inserting cart")
			return insertCart, errInsert
		}
	}

	return insertCart, nil
}

func (r *Repository_Cart) GetCart(id int) ([]entities.Cart, error) {
	var allCarts []entities.Cart

	result := r.db.Where("user_id = ?", id).Find(&allCarts)

	err := result.Error

	if err != nil {
		return allCarts, err
	}

	return allCarts, nil
}

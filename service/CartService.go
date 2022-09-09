package service

import (
	"efishery-project/entities"
	"efishery-project/repository"
)

type ServiceCart interface {
	InsertToCartService(cart entities.Cart) (entities.Cart, error)
	GetCartService(id int) ([]entities.Cart, error)
}

type Service_Cart struct {
	repositorycart repository.RepositoryCart
}

func NewServiceCart(repositorycart repository.RepositoryCart) *Service_Cart {
	return &Service_Cart{repositorycart}
}

func (s *Service_Cart) InsertToCartService(cart entities.Cart) (entities.Cart, error) {
	insertToCart, errInsertToCart := s.repositorycart.InsertToCart(cart)

	if errInsertToCart != nil {
		return insertToCart, errInsertToCart
	}

	return insertToCart, nil
}

func (s *Service_Cart) GetCartService(id int) ([]entities.Cart, error) {
	getCart, errGetCart := s.repositorycart.GetCart(id)

	if errGetCart != nil {
		return getCart, errGetCart
	}

	return getCart, nil
}

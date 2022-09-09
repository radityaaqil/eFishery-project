package service

import (
	"efishery-project/entities"
	"efishery-project/repository"
)

type ServiceProduct interface {
	GetListProductService(page int, category string, minPrice int, maxPrice int) ([]entities.ProductList, error)
	GetProductDetailService(id int) (entities.ProductDetail, error)
}

type Service_Product struct {
	repositoryproduct repository.RepositoryProduct
}

func NewServiceProduct(repositoryproduct repository.RepositoryProduct) *Service_Product {
	return &Service_Product{repositoryproduct}
}

func (s *Service_Product) GetListProductService(page int, category string, minPrice int, maxPrice int) ([]entities.ProductList, error) {
	productList, errProductList := s.repositoryproduct.GetListProduct(page, category, minPrice, maxPrice)

	if errProductList != nil {
		return productList, errProductList
	}

	return productList, nil
}

func (s *Service_Product) GetProductDetailService(id int) (entities.ProductDetail, error) {
	productDetail, errProductDetail := s.repositoryproduct.GetProductDetail(id)

	if errProductDetail != nil {
		return productDetail, errProductDetail
	}

	return productDetail, nil
}

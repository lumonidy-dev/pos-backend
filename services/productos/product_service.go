package services

import (
	"pos-backend/entity"
)

// ProductService define las operaciones que se pueden realizar con los productos
type ProductService interface {
	CreateProduct(product *entity.Productos) (*entity.Productos, error)
	GetAllProducts() ([]*entity.Productos, error)
	GetProductByID(id string) (*entity.Productos, error)
	UpdateProduct(product *entity.Productos, docID string) (*entity.Productos, error)
	DeleteProduct(id string) error
}

// Path: services/product_service.go

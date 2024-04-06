package services

import (
	"context"
	"pos-backend/entity"

	"cloud.google.com/go/firestore"
)

// ProductService define las operaciones que se pueden realizar con los productos
type ProductService interface {
	CreateProduct(product *entity.Productos) (*entity.Productos, error)
	GetAllProducts() ([]*entity.ProductosResponse, error)
	GetProductByID(id string) (*entity.ProductosResponse, error)
	UpdateProduct(product *entity.Productos, docID string) (*entity.Productos, error)
	DeleteProduct(id string) error
}

// ProductServiceFirestore es una implementación de ProductService que utiliza Firestore como base de datos
type ProductServiceFirestore struct {
	client *firestore.Client
	ctx    context.Context
}

// NewProductServiceFirestore crea una nueva instancia de ProductServiceFirestore
func NewProductServiceFirestore(client *firestore.Client, ctx context.Context) *ProductServiceFirestore {
	return &ProductServiceFirestore{
		client: client,
		ctx:    ctx,
	}
}

// Path: services/productos/service.go

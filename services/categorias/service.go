package services

import (
	"context"
	"pos-backend/entity"

	"cloud.google.com/go/firestore"
)

// CategoryService define las operaciones que se pueden realizar con las categorías
type CategoryService interface {
	CreateCategory(category *entity.Categorias) (*entity.Categorias, error)
	GetAllCategories() ([]*entity.CategoriasResponse, error)
	GetCategoryByID(id string) (*entity.CategoriasResponse, error)
	UpdateCategory(category *entity.Categorias, docID string) (*entity.Categorias, error)
	DeleteCategory(id string) error
}

type CategoryServiceFirestore struct {
	client *firestore.Client
	ctx    context.Context
}

func NewCategoryServiceFirestore(client *firestore.Client, ctx context.Context) *CategoryServiceFirestore {
	return &CategoryServiceFirestore{
		client: client,
		ctx:    ctx,
	}
}

// Path: services/categorias/service.go

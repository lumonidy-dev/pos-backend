package services

import (
	"context"
	"log"
	"pos-backend/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

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

func (s *ProductServiceFirestore) CreateProduct(product *entity.Productos) (*entity.Productos, error) {
	_, _, err := s.client.Collection("Productos").Add(s.ctx, product)
	if err != nil {
		log.Printf("Error al añadir un producto: %v", err)
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceFirestore) GetAllProducts() ([]*entity.Productos, error) {
	iter := s.client.Collection("Productos").Documents(s.ctx)
	var products []*entity.Productos
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error al obtener los productos: %v", err)
			return nil, err
		}
		var product entity.Productos
		if err := doc.DataTo(&product); err != nil {
			log.Printf("Error al convertir datos a struct: %v", err)
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (s *ProductServiceFirestore) GetProductByID(docID string) (*entity.Productos, error) {
	doc, err := s.client.Collection("ss").Doc(docID).Get(s.ctx)
	if err != nil {
		log.Printf("Error al obtener el producto por ID: %v", err)
		return nil, err
	}
	var product entity.Productos
	if err := doc.DataTo(&product); err != nil {
		log.Printf("Error al convertir datos a struct: %v", err)
		return nil, err
	}
	return &product, nil
}

func (s *ProductServiceFirestore) UpdateProduct(product *entity.Productos, docID string) (*entity.Productos, error) {
	// Utilizar el ID del documento en la consulta
	_, err := s.client.Collection("Productos").Doc(docID).Set(s.ctx, product)
	if err != nil {
		log.Printf("Error al actualizar un producto: %v", err)
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceFirestore) DeleteProduct(id string) error {
	_, err := s.client.Collection("Productos").Doc(id).Delete(s.ctx)
	if err != nil {
		log.Printf("Error al eliminar un producto: %v", err)
		return err
	}
	return nil
}

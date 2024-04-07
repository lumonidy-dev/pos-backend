package services

import (
	"log"
	"pos-backend/entity"

	"google.golang.org/api/iterator"
)

// GetAllProducts obtiene todos los productos junto con sus IDs de documento
func (s *ProductServiceFirestore) GetAllProducts() ([]*entity.ProductosResponse, error) {
	iter := s.client.Collection("Productos").Documents(s.ctx)
	var products []*entity.ProductosResponse
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

		products = append(products, &entity.ProductosResponse{
			ID:           doc.Ref.ID,
			Nombre:       product.Nombre,
			Precio:       product.Precio,
			Ingredientes: product.Ingredientes,
			Stock:        product.Stock,
			Categoria:    product.Categoria,
		})
	}
	return products, nil
}

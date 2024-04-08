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

		// Obtener el nombre de la categoría
		categoryDoc, err := s.client.Collection("Categorias").Doc(product.Categoria).Get(s.ctx)
		if err != nil {
			log.Printf("Error al obtener la categoría: %v", err)
			continue
		}

		categoryData := categoryDoc.Data()
		categoryName, ok := categoryData["nombre"].(string)
		if !ok {
			log.Printf("El campo Nombre no es de tipo string")
			return nil, err
		}

		products = append(products, &entity.ProductosResponse{
			ID:           doc.Ref.ID,
			Nombre:       product.Nombre,
			Precio:       product.Precio,
			Ingredientes: product.Ingredientes,
			Stock:        product.Stock,
			Categoria: entity.Categoria{
				ID:     categoryDoc.Ref.ID,
				Nombre: categoryName,
			},
		})
	}
	return products, nil
}

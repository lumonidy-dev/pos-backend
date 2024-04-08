package services

import (
	"log"
	"pos-backend/entity"
)

func (s *ProductServiceFirestore) GetProductByID(docID string) (*entity.ProductosResponse, error) {
	doc, err := s.client.Collection("Productos").Doc(docID).Get(s.ctx)
	if err != nil {
		log.Printf("Error al obtener el producto: %v", err)
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
		// Si no se puede obtener la categoría, no se menciona en la respuesta
		return &entity.ProductosResponse{
			ID:           doc.Ref.ID,
			Nombre:       product.Nombre,
			Precio:       product.Precio,
			Ingredientes: product.Ingredientes,
			Stock:        product.Stock,
		}, nil
	}

	categoryData := categoryDoc.Data()
	categoryName, ok := categoryData["nombre"].(string)
	if !ok {
		log.Printf("El campo Nombre no es de tipo string")
		return nil, err
	}

	return &entity.ProductosResponse{
		ID:           doc.Ref.ID,
		Nombre:       product.Nombre,
		Precio:       product.Precio,
		Ingredientes: product.Ingredientes,
		Stock:        product.Stock,
		Categoria: entity.Categoria{
			ID:     product.Categoria,
			Nombre: categoryName,
		},
	}, nil
}

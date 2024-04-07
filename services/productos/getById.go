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
	return &entity.ProductosResponse{
		ID:           doc.Ref.ID,
		Nombre:       product.Nombre,
		Precio:       product.Precio,
		Ingredientes: product.Ingredientes,
		Stock:        product.Stock,
		Categoria:    product.Categoria,
	}, nil
}

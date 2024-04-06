package services

import (
	"log"
	"pos-backend/entity"
)

func (s *ProductServiceFirestore) CreateProduct(product *entity.Productos) (*entity.Productos, error) {
	_, _, err := s.client.Collection("Productos").Add(s.ctx, product)
	if err != nil {
		log.Printf("Error al añadir un producto: %v", err)
		return nil, err
	}
	return product, nil
}

package services

import (
	"log"
	"pos-backend/entity"
)

func (s *ProductServiceFirestore) UpdateProduct(product *entity.Productos, docID string) (*entity.Productos, error) {
	// Utilizar el ID del documento en la consulta
	_, err := s.client.Collection("Productos").Doc(docID).Set(s.ctx, product)
	if err != nil {
		log.Printf("Error al actualizar un producto: %v", err)
		return nil, err
	}
	return product, nil
}

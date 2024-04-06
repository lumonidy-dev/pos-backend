package services

import (
	"log"
	"pos-backend/entity"
)

func (s *CategoryServiceFirestore) UpdateCategory(category *entity.Categorias, docID string) (*entity.Categorias, error) {
	// Utilizar el ID del documento en la consulta
	_, err := s.client.Collection("Categorias").Doc(docID).Set(s.ctx, category)
	if err != nil {
		log.Printf("Error al actualizar una categoria: %v", err)
		return nil, err
	}
	return category, nil
}

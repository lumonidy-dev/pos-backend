package services

import (
	"log"
	"pos-backend/entity"
)

func (s *CategoryServiceFirestore) CreateCategory(category *entity.Categorias) (*entity.Categorias, error) {
	_, _, err := s.client.Collection("Categorias").Add(s.ctx, category)
	if err != nil {
		log.Printf("Error al añadir una categoria: %v", err)
		return nil, err
	}
	return category, nil
}

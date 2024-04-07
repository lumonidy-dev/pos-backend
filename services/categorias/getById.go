package services

import (
	"log"
	"pos-backend/entity"
)

func (s *CategoryServiceFirestore) GetCategoryByID(docID string) (*entity.CategoriasResponse, error) {
	doc, err := s.client.Collection("Categorias").Doc(docID).Get(s.ctx)
	if err != nil {
		log.Printf("Error al obtener una categoria por ID: %v", err)
		return nil, err
	}

	var category entity.Categorias
	if err := doc.DataTo(&category); err != nil {
		log.Printf("Error al convertir datos a struct: %v", err)
		return nil, err
	}

	return &entity.CategoriasResponse{
		ID:     doc.Ref.ID,
		Nombre: category.Nombre,
		Desc:   category.Desc,
		URL:    category.URL,
		Tipos:  category.Tipos,
	}, nil
}

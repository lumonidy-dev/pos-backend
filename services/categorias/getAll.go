package services

import (
	"log"
	"pos-backend/entity"

	"google.golang.org/api/iterator"
)

func (s *CategoryServiceFirestore) GetAllCategories() ([]*entity.CategoriasResponse, error) {
	iter := s.client.Collection("Categorias").Documents(s.ctx)
	var categories []*entity.CategoriasResponse

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error al obtener las categorías: %v", err)
			return nil, err
		}

		var category entity.Categorias
		if err := doc.DataTo(&category); err != nil {
			log.Printf("Error al convertir datos a struct: %v", err)
			return nil, err
		}

		categories = append(categories, &entity.CategoriasResponse{
			ID:     doc.Ref.ID,
			Nombre: category.Nombre,
			Desc:   category.Desc,
			URL:    category.URL,
			Tipos:  category.Tipos,
		})
	}

	return categories, nil
}

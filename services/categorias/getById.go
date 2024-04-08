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

	// Obtener los nombres de los productos
	// category.Tipos es un slice de IDs de productos de la colección Productos
	// Se debe obtener el nombre de cada producto para mostrarlo en la respuesta

	// Crear un slice de TipoProducto para almacenar los nombres de los productos
	var tipos []entity.TipoProducto

	for _, tipoID := range category.Tipos {
		tipoDoc, err := s.client.Collection("Productos").Doc(tipoID).Get(s.ctx)
		if err != nil {
			log.Printf("Error al obtener el producto: %v", err)
			continue
		}

		var tipo entity.Productos
		if err := tipoDoc.DataTo(&tipo); err != nil {
			log.Printf("Error al convertir datos a struct: %v", err)
			return nil, err
		}

		tipos = append(tipos, entity.TipoProducto{
			ID:           tipoDoc.Ref.ID,
			Nombre:       tipo.Nombre,
			Precio:       tipo.Precio,
			Stock:        tipo.Stock,
			Ingredientes: tipo.Ingredientes,
		})
	}

	return &entity.CategoriasResponse{
		ID:     doc.Ref.ID,
		Nombre: category.Nombre,
		Desc:   category.Desc,
		URL:    category.URL,
		Tipos:  tipos,
	}, nil
}

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

		// Obtener los nombres de los productos
		// category.Tipos es un slice de IDs de productos de la colección Productos
		// Se debe obtener el nombre de cada producto para mostrarlo en la respuesta

		var tipos []entity.TipoProducto // Crear un slice de TipoProducto para almacenar los nombres de los productos

		for _, tipoID := range category.Tipos {
			if tipoID == "" {
				// Si el ID del producto está vacío, se agrega un producto falso
				tipos = append(tipos, entity.TipoProducto{
					ID:           tipoID,                   // ID para "Sin categoría"
					Nombre:       "Producto no encontrado", // Nombre predeterminado
					Precio:       "0",                      // Precio predeterminado
					Stock:        "0",                      // Stock predeterminado
					Ingredientes: []string{},               // Lista de ingredientes vacía
				})
				continue
			}

			// Consultar el producto en la base de datos
			tipoDoc, err := s.client.Collection("Productos").Doc(tipoID).Get(s.ctx)
			if err != nil {
				log.Printf("Error al obtener el producto: %v", err)
				// Si hay un error al obtener el producto, se crea un producto falso
				tipos = append(tipos, entity.TipoProducto{
					ID:           tipoID,                   // ID para "Sin categoría"
					Nombre:       "Producto no encontrado", // Nombre predeterminado
					Precio:       "0",                      // Precio predeterminado
					Stock:        "0",                      // Stock predeterminado
					Ingredientes: []string{},               // Lista de ingredientes vacía
				})
				continue
			}

			var tipo entity.Productos
			if err := tipoDoc.DataTo(&tipo); err != nil {
				log.Printf("Error al convertir datos a struct: %v", err)
				return nil, err
			} // Si no puede convertir los datos, se le agregará un valor predeterminado

			tipos = append(tipos, entity.TipoProducto{
				ID:           tipoDoc.Ref.ID,
				Nombre:       tipo.Nombre,
				Precio:       tipo.Precio,
				Stock:        tipo.Stock,
				Ingredientes: tipo.Ingredientes,
			})
		}

		categories = append(categories, &entity.CategoriasResponse{
			ID:     doc.Ref.ID,
			Nombre: category.Nombre,
			Desc:   category.Desc,
			URL:    category.URL,
			Tipos:  tipos,
		})
	}

	return categories, nil
}

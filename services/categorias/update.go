package services

import (
	"log"
	"pos-backend/entity"

	"cloud.google.com/go/firestore"
)

func (s *CategoryServiceFirestore) UpdateCategory(category *entity.Categorias, docID string) (*entity.Categorias, error) {
	// Obtener la categoría actual
	oldCategoryDoc, err := s.client.Collection("Categorias").Doc(docID).Get(s.ctx)
	if err != nil {
		log.Printf("Error al obtener el documento de la categoría: %v", err)
		return nil, err
	}
	oldCategoryData := oldCategoryDoc.Data()

	// Crear un mapa para almacenar los productos que están en la categoría actual
	oldProducts := make(map[string]bool)
	if oldCategoryData != nil {
		tiposInterface := oldCategoryData["tipos"]
		if tiposInterface != nil {
			tiposSlice := tiposInterface.([]interface{})
			for _, tipo := range tiposSlice {
				oldProducts[tipo.(string)] = true
			}
		}
	}

	// Iterar sobre los productos de la nueva categoría
	for _, productID := range category.Tipos {
		// Verificar si el producto ya estaba en la categoría anterior
		if oldProducts[productID] {
			// El producto ya estaba en la categoría anterior, no es necesario hacer nada
			delete(oldProducts, productID)
		} else {
			// El producto no estaba en la categoría anterior, asignarle esta categoría
			if err := s.updateProductCategory(productID, docID); err != nil {
				log.Printf("Error al actualizar la categoría del producto: %v", err)
				// Puedes manejar el error según sea necesario (p. ej., registrar, devolver el error, etc.)
			}
		}
	}

	// Iterar sobre los productos que estaban en la categoría anterior pero no están en la nueva
	for productID := range oldProducts {
		// Asignar la categoría "Sin categoría" a estos productos
		if err := s.updateProductCategory(productID, "UfsOblViwVNpqB1ycVJG"); err != nil {
			log.Printf("Error al actualizar la categoría del producto: %v", err)
			// Puedes manejar el error según sea necesario (p. ej., registrar, devolver el error, etc.)
		}
	}

	// Actualizar la categoría en Firestore
	_, err = s.client.Collection("Categorias").Doc(docID).Set(s.ctx, category)
	if err != nil {
		log.Printf("Error al actualizar una categoria: %v", err)
		return nil, err
	}

	return category, nil
}

func (s *CategoryServiceFirestore) updateProductCategory(productID, newCategoryID string) error {
	// Actualizar la categoría del producto
	_, err := s.client.Collection("Productos").Doc(productID).Update(s.ctx, []firestore.Update{
		{Path: "categoria", Value: newCategoryID},
	})
	return err
}

package services

import (
	"log"
	"pos-backend/entity"

	"cloud.google.com/go/firestore"
)

func (s *ProductServiceFirestore) UpdateProduct(product *entity.Productos, docID string) (*entity.Productos, error) {

	// Revisar el campo [categoria: string] que corresponde a una supuesta ID de categoría
	// si la categoría no existe, se crea una nueva
	// si la categoría existe, se busca en ella, en el campo [tipos: array] si el producto ya existe
	// si el producto no existe, se agrega a la categoría
	// si el producto ya existe, se actualiza
	// si el producto ya existe, se elimina de la categoría anterior
	// si la categoría anterior queda vacía, se elimina

	// Obtener el ID de la categoría del producto actual
	oldCategoryID := ""
	oldProductDoc, err := s.client.Collection("Productos").Doc(docID).Get(s.ctx)
	if err != nil {
		log.Printf("Error al obtener el documento del producto: %v", err)
		return nil, err
	}
	oldProductData := oldProductDoc.Data()
	if oldProductData != nil {
		oldCategoryID = oldProductData["categoria"].(string)
	}
	// Si la categoría del producto ha cambiado
	if product.Categoria != oldCategoryID {
		// Si hay una categoría anterior, eliminar el producto de su lista de tipos
		if oldCategoryID != "" {
			// Eliminar el producto del array de tipos en la categoría anterior
			if err := s.removeProductFromCategory(oldCategoryID, docID); err != nil {
				log.Printf("Error al eliminar el producto de la categoría anterior: %v", err)
				// Puedes manejar el error según sea necesario (p. ej., registrar, devolver el error, etc.)
			}
		}

		// Agregar el producto a la nueva categoría
		if err := s.addProductToCategory(product.Categoria, docID); err != nil {
			log.Printf("Error al agregar el producto a la nueva categoría: %v", err)
			// Puedes manejar el error según sea necesario (p. ej., registrar, devolver el error, etc.)
		}
	}

	// Actualizar el producto en Firestore
	_, err = s.client.Collection("Productos").Doc(docID).Set(s.ctx, product)
	if err != nil {
		log.Printf("Error al actualizar un producto: %v", err)
		return nil, err
	}

	return product, nil
}

func (s *ProductServiceFirestore) removeProductFromCategory(categoryID, productID string) error {
	// Eliminar el producto del array de tipos en la categoría anterior
	_, err := s.client.Collection("Categorias").Doc(categoryID).Update(s.ctx, []firestore.Update{
		{Path: "tipos", Value: firestore.ArrayRemove(productID)},
	})
	return err
}

func (s *ProductServiceFirestore) addProductToCategory(categoryID, productID string) error {
	// Agregar el producto al array de tipos en la nueva categoría
	_, err := s.client.Collection("Categorias").Doc(categoryID).Update(s.ctx, []firestore.Update{
		{Path: "tipos", Value: firestore.ArrayUnion(productID)},
	})
	return err
}

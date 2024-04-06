package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/productos"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context" // Importar el paquete de contexto
)

// GetProductByIDHandler maneja las solicitudes para obtener un producto por su ID
func GetProductByIDHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el producto por su ID utilizando el servicio
		productID := r.URL.Query().Get("id")
		if productID == "" {
			http.Error(w, "Se requiere el ID del producto", http.StatusBadRequest)
			return
		}

		product, err := productService.GetProductByID(productID)
		if err != nil {
			http.Error(w, "Error al obtener el producto", http.StatusInternalServerError)
			return
		}

		// Obtener el nombre de la categoría
		nombreCategoria, err := ObtenerNombreCategoria(product.Categoria, r.Context())
		if err != nil {
			http.Error(w, "Error al obtener el nombre de la categoría", http.StatusInternalServerError)
			return
		}

		// Crear una instancia de ProductosResponse y asignar los valores
		response := entity.ProductosResponse{
			ID:           productID, // Asignar manualmente el ID
			Nombre:       product.Nombre,
			Precio:       product.Precio,
			Ingredientes: product.Ingredientes,
			Stock:        product.Stock,
			Categoria:    nombreCategoria, // Asignar el nombre de la categoría
		}

		// Responder con la instancia de ProductosResponse
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// ObtenerNombreCategoria obtiene el nombre de la categoría a partir de la referencia del documento
func ObtenerNombreCategoria(categoriaRef *firestore.DocumentRef, ctx context.Context) (string, error) {
	// Obtener el documento de la categoría utilizando la referencia del documento
	categoriaDoc, err := categoriaRef.Get(ctx)
	if err != nil {
		return "", err
	}

	// Obtener el campo "nombre" del documento de la categoría
	nombre, ok := categoriaDoc.Data()["nombre"].(string)
	if !ok {
		return "", fmt.Errorf("no se encontró el nombre de la categoría")
	}

	return nombre, nil
}

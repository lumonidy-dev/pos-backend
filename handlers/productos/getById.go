package handlers

import (
	"encoding/json"
	"net/http"
	productos "pos-backend/services/productos"

	"github.com/gorilla/mux"
)

// GetProductByIDHandler maneja las solicitudes para obtener un producto por su ID
func GetProductByIDHandler(productService productos.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del producto de los parámetros de la URL
		vars := mux.Vars(r)
		productID := vars["id"]
		if productID == "" {
			http.Error(w, "Se requiere el ID del producto", http.StatusBadRequest)
			return
		}

		// Obtener el producto por su ID utilizando el servicio
		product, err := productService.GetProductByID(productID)
		if err != nil {
			http.Error(w, "Error al obtener el producto", http.StatusInternalServerError)
			return
		}

		// Responder con el producto encontrado
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
	services "pos-backend/services/productos"
)

// GetAllProductsHandler maneja las solicitudes para obtener todos los productos
func GetAllProductsHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := productService.GetAllProducts()
		if err != nil {
			http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

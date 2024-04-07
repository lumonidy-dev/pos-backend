package handlers

import (
	"encoding/json"
	"net/http"
	services "pos-backend/services/productos"
)

// GetAllProductsHandler maneja las solicitudes para obtener todos los productos
// @Summary Obtener todos los productos
// @Description Obtiene todos los productos registrados en el sistema
// @Tags Productos
// @Produce json
// @Success 200 {array} entity.Productos "Lista de productos"
// @Failure 500 "Error al obtener los productos"
// @Router /productos [get]
func GetAllProductsHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := productService.GetAllProducts()
		if err != nil {
			http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

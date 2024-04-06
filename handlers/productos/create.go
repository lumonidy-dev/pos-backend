package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/productos"
)

// CreateProductHandler maneja las solicitudes para crear un nuevo producto
func CreateProductHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product entity.Productos
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "Error decodificando el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		createdProduct, err := productService.CreateProduct(&product)
		if err != nil {
			http.Error(w, "Error al crear el producto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdProduct)
	}
}

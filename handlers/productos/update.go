package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/productos"
)

func UpdateProductHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del producto de los parámetros de la URL
		productID := r.URL.Query().Get("id")
		if productID == "" {
			http.Error(w, "Se requiere el ID del producto", http.StatusBadRequest)
			return
		}

		// Decodificar el cuerpo de la solicitud para obtener los nuevos datos del producto
		var updatedProduct entity.Productos
		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, "Error decodificando el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		// Actualizar el producto utilizando el servicio
		_, err = productService.UpdateProduct(&updatedProduct, productID)
		if err != nil {
			http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
			return
		}

		// Responder con el producto actualizado
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedProduct)
	}
}

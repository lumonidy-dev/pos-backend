package handlers

import (
	"net/http"
	services "pos-backend/services/productos"
)

// DeleteProductHandler maneja las solicitudes para eliminar un producto
func DeleteProductHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del producto de los parámetros de la URL
		productID := r.URL.Query().Get("id")
		if productID == "" {
			http.Error(w, "Se requiere el ID del producto", http.StatusBadRequest)
			return
		}

		// Eliminar el producto utilizando el servicio
		err := productService.DeleteProduct(productID)
		if err != nil {
			http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

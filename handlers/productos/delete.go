package handlers

import (
	"net/http"
	services "pos-backend/services/productos"

	"github.com/gorilla/mux"
)

// DeleteProductHandler maneja las solicitudes para eliminar un producto
// @Summary Eliminar un producto
// @Description Elimina un producto del sistema
// @Tags Productos
// @Param id query string true "ID del producto a eliminar"
// @Success 200 "Producto eliminado exitosamente"
// @Failure 400 "Se requiere el ID del producto"
// @Failure 500 "Error al eliminar el producto"
// @Router /productos [delete]
func DeleteProductHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtiene el ID de la categoría de los parámetros de la URL
		vars := mux.Vars(r)
		productID := vars["id"]
		if productID == "" {
			http.Error(w, "Se requiere el ID de la categoría en la URL", http.StatusBadRequest)
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

package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/productos"

	"github.com/gorilla/mux"
)

// UpdateProductHandler maneja las solicitudes para actualizar un producto
// @Summary Actualizar un producto
// @Description Actualiza un producto en el sistema
// @Tags Productos
// @Param id query string true "ID del producto a actualizar"
// @Accept json
// @Produce json
// @Param body body entity.Productos true "Cuerpo de la solicitud en formato JSON con los datos actualizados del producto"
// @Success 200 {object} entity.Productos "Producto actualizado exitosamente"
// @Failure 400 {string} string "Se requiere el ID del producto o el cuerpo de la solicitud está mal formado"
// @Failure 500 {string} string "Error al actualizar el producto"
// @Router /productos [put]
func UpdateProductHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtiene el ID de la categoría de los parámetros de la URL
		vars := mux.Vars(r)
		productID := vars["id"]
		if productID == "" {
			http.Error(w, "Se requiere el ID de la categoría en la URL", http.StatusBadRequest)
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

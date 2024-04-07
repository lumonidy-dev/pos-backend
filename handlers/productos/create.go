package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/productos"
)

// CreateProductHandler maneja las solicitudes para crear un nuevo producto
// @Summary Crear un nuevo producto
// @Description Crea un nuevo producto en el sistema
// @Tags Productos
// @Accept json
// @Produce json
// @Param body body entity.Productos true "Cuerpo de la solicitud en formato JSON con los datos del nuevo producto"
// @Success 201 {object} entity.Productos "Producto creado exitosamente"
// @Failure 400 {string} string "El cuerpo de la solicitud está mal formado"
// @Failure 500 {string} string "Error al crear el producto"
// @Router /productos [post]
func CreateProductHandler(productService services.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product entity.Productos
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "El cuerpo de la solicitud está mal formado", http.StatusBadRequest)
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

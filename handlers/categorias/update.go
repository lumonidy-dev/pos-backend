package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/categorias"
)

// UpdateCategoryHandler maneja las solicitudes para actualizar una categoría
func UpdateCategoryHandler(categoryService services.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtiene el ID de la categoría de los parámetros de la URL
		categoryID := r.URL.Query().Get("id")
		if categoryID == "" {
			http.Error(w, "Se requiere el ID de la categoría", http.StatusBadRequest)
			return
		}

		// Decodifica el cuerpo de la solicitud para obtener los nuevos datos de la categoría
		var updatedCategory entity.Categorias
		err := json.NewDecoder(r.Body).Decode(&updatedCategory)
		if err != nil {
			http.Error(w, "Error decodificando el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		// Actualiza la categoría utilizando el servicio
		_, err = categoryService.UpdateCategory(&updatedCategory, categoryID)

	}
}

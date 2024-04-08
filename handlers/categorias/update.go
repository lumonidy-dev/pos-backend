package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/categorias"

	"github.com/gorilla/mux"
)

// @Summary Actualizar una categoría existente
// @Description Actualiza los datos de una categoría existente en el sistema
// @Tags Categorías
// @Accept json
// @Produce json
// @Param id query string true "ID de la categoría a actualizar"
// @Param body body entity.Categorias true "Cuerpo de la solicitud en formato JSON con los datos actualizados de la categoría"
// @Success 200 {string} string "Categoría actualizada exitosamente"
// @Failure 400 {string} string "Se requiere el ID de la categoría o el cuerpo de la solicitud está mal formado"
// @Failure 500 {string} string "Error al actualizar la categoría en el sistema"
func UpdateCategoryHandler(categoryService services.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtiene el ID de la categoría de los parámetros de la URL
		vars := mux.Vars(r)
		categoryID := vars["id"]
		if categoryID == "" {
			http.Error(w, "Se requiere el ID de la categoría en la URL", http.StatusBadRequest)
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
		if err != nil {
			http.Error(w, "Error al actualizar la categoría", http.StatusInternalServerError)
			return
		}
		// La categoría se actualizó exitosamente
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Categoría actualizada exitosamente"))
	}
}

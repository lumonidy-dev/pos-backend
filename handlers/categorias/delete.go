package handlers

import (
	"net/http"
	categorias "pos-backend/services/categorias"

	"github.com/gorilla/mux"
)

// DeleteCategoryHandler maneja las solicitudes para eliminar una categoría
// @Summary Eliminar una categoría
// @Description Elimina una categoría del sistema
// @Tags Categorías
// @Param id query string true "ID de la categoría a eliminar"
// @Success 200 "Categoría eliminada"
// @Failure 400 "Se requiere el ID de la categoría"
// @Failure 500 "Error al eliminar la categoría"
// @Router /categorias [delete]
func DeleteCategoryHandler(categoryService categorias.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtiene el ID de la categoría de los parámetros de la URL
		vars := mux.Vars(r)
		categoryID := vars["id"]
		if categoryID == "" {
			http.Error(w, "Se requiere el ID de la categoría en la URL", http.StatusBadRequest)
			return
		}

		// Elimina la categoría utilizando el servicio
		err := categoryService.DeleteCategory(categoryID)
		if err != nil {
			http.Error(w, "Error al eliminar la categoría", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

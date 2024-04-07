package handlers

import (
	"encoding/json"
	"net/http"
	services "pos-backend/services/categorias"
)

// GetAllCategoriesHandler maneja las solicitudes para obtener todas las categorías
// @Summary Obtener todas las categorías
// @Description Obtiene todas las categorías del sistema
// @Tags Categorías
// @Produce json
// @Success 200 {array} entity.CategoriasResponse
// @Failure 500 "Error al obtener las categorías"
// @Router /categorias [get]
func GetAllCategoriesHandler(categoryService services.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		categories, err := categoryService.GetAllCategories()
		if err != nil {
			http.Error(w, "Error al obtener las categorías", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(categories)
	}
}

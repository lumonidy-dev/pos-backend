package handlers

import (
	"encoding/json"
	"net/http"
	services "pos-backend/services/categorias"
)

// GetAllCategoriesHandler maneja las solicitudes para obtener todas las categorías
func GetAllCategoriesHandler(categoryService services.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := categoryService.GetAllCategories()
		if err != nil {
			http.Error(w, "Error al obtener las categorías", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	}
}

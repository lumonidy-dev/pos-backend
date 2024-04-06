package handlers

import (
	"net/http"
	services "pos-backend/services/categorias"
)

func DeleteCategoryHandler(categoryService services.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryID := r.URL.Query().Get("id")
		if categoryID == "" {
			http.Error(w, "Se requiere el ID de la categoría", http.StatusBadRequest)
			return
		}

		err := categoryService.DeleteCategory(categoryID)
		if err != nil {
			http.Error(w, "Error al eliminar la categoría", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

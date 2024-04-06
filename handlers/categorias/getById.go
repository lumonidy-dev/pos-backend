package handlers

import (
	"encoding/json"
	"net/http"
	categorias "pos-backend/services/categorias"

	"github.com/gorilla/mux"
)

// GetByIDCategoryHandler maneja las solicitudes para obtener una categoría por su ID
func GetCategoryByIDHandler(categoryService categorias.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Otener el ID de la categoría de los parámetros de la URL
		vars := mux.Vars(r)
		categoryID := vars["id"]
		if categoryID == "" {
			http.Error(w, "Se requiere el ID de la categoría", http.StatusBadRequest)
			return
		}

		// Obtener la categoría por su ID utilizando el servicio
		category, err := categoryService.GetCategoryByID(categoryID)
		if err != nil {
			http.Error(w, "Error al obtener la categoría", http.StatusInternalServerError)
			return
		}

		// Responder con la categoría encontrada
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(category)
	}
}

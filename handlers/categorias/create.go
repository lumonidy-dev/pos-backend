package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/categorias"
)

// CreateCategoryHandler maneja las solicitudes para crear una nueva categoría
func CreateCategoryHandler(categoryService services.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category entity.Categorias
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			http.Error(w, "Error decodificando el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		createdCategory, err := categoryService.CreateCategory(&category)
		if err != nil {
			http.Error(w, "Error al crear la categoría", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdCategory)
	}
}

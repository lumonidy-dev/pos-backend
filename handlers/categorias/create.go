package handlers

import (
	"encoding/json"
	"net/http"
	"pos-backend/entity"
	services "pos-backend/services/categorias"
)

// @Summary Crear una nueva categoría
// @Description Crea una nueva categoría en el sistema
// @Tags Categorías
// @Accept json
// @Produce json
// @Param body body entity.Categorias true "Cuerpo de la solicitud en formato JSON"
// @Success 201 {object} entity.CategoriasResponse
// @Router /categorias [post]
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

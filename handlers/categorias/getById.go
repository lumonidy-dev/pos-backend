package handlers

import (
	"encoding/json"
	"net/http"
	categorias "pos-backend/services/categorias"

	"github.com/gorilla/mux"
)

// GetByIDCategoryHandler maneja las solicitudes para obtener una categoría por su ID
// @Summary Obtener una categoría por su ID
// @Description Obtiene una categoría del sistema por su ID
// @Tags Categorías
// @Param id path string true "ID de la categoría a obtener"
// @Produce json
// @Success 200 {object} entity.CategoriasResponse
// @Failure 400 "Se requiere el ID de la categoría"
// @Failure 500 "Error al obtener la categoría"
// @Router /categorias/{id} [get]
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(category)
	}
}

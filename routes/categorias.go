package routes

import (
	handlers "pos-backend/handlers/categorias"
	services "pos-backend/services/categorias"

	"github.com/gorilla/mux"
)

// SetCategoriasRoutes configura las rutas para las categorías
func SetCategoriasRoutes(router *mux.Router, categoryService services.CategoryService) {
	router.HandleFunc("/categorias", handlers.CreateCategoryHandler(categoryService)).Methods("POST")
	router.HandleFunc("/categorias", handlers.GetAllCategoriesHandler(categoryService)).Methods("GET")
	router.HandleFunc("/categorias/{id}", handlers.GetCategoryByIDHandler(categoryService)).Methods("GET")
	router.HandleFunc("/categorias/{id}", handlers.UpdateCategoryHandler(categoryService)).Methods("PUT")
	router.HandleFunc("/categorias/{id}", handlers.DeleteCategoryHandler(categoryService)).Methods("DELETE")
}

// Path: routes/categorias.go

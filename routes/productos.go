package routes

import (
	handlers "pos-backend/handlers/productos"
	services "pos-backend/services/productos"

	"github.com/gorilla/mux"
)

func SetProductosRoutes(router *mux.Router, productService services.ProductService) {
	router.HandleFunc("/productos", handlers.CreateProductHandler(productService)).Methods("POST")
	router.HandleFunc("/productos", handlers.GetAllProductsHandler(productService)).Methods("GET")
	router.HandleFunc("/productos/{id}", handlers.GetProductByIDHandler(productService)).Methods("GET")
	router.HandleFunc("/productos/{id}", handlers.UpdateProductHandler(productService)).Methods("PUT")
	router.HandleFunc("/productos/{id}", handlers.DeleteProductHandler(productService)).Methods("DELETE")
}

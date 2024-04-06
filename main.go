package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pos-backend/middleware"
	"pos-backend/routes"
	categorias "pos-backend/services/categorias"
	productos "pos-backend/services/productos"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var (
	credentialPath = "credentials.json"
	projectId      = "lumo-pos"
)

func main() {
	// Inicializar la app de Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialPath)
	app, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Error al crear el cliente de Firestore: %v", err)
	}
	defer app.Close() // Cerrar el cliente de Firestore al salir de main()

	fmt.Println("Firebase inicializado correctamente")

	// Configurar enrutador
	router := mux.NewRouter()

	// Configurar middleware para registrar la IP del cliente en todos los manejadores de rutas
	router.Use(middleware.LogIPMiddleware)

	// Configurar productos service
	productService := productos.NewProductServiceFirestore(app, ctx)

	// Configurar categorías service
	categoryService := categorias.NewCategoryServiceFirestore(app, ctx)
	// Configurar rutas
	routes.SetCategoriasRoutes(router, categoryService)
	routes.SetProductosRoutes(router, productService)

	// Configurar ruta de prueba
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Up and running...")
	})

	// Configurar ruta de documentación
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	// Configurar servidor HTTP
	const port = ":8000"
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Iniciar servidor HTTP
	log.Println("Server running on port", port)
	log.Fatalln(server.ListenAndServe())
}

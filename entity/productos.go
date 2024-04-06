package entity

import "cloud.google.com/go/firestore"

type Productos struct {
	Nombre       string                 `firestore:"nombre"`
	Precio       string                 `firestore:"precio"`
	Ingredientes []string               `firestore:"ingredientes"`
	Stock        int                    `firestore:"stock"`
	Categoria    *firestore.DocumentRef `firestore:"categoria"`
}
type ProductosResponse struct {
	ID           string   `json:"id"`
	Nombre       string   `json:"nombre"`
	Precio       string   `json:"precio"`
	Ingredientes []string `json:"ingredientes"`
	Stock        int      `json:"stock"`
	Categoria    string   `json:"categoria"`
}

// Path: entity/categorias.go

package entity

type Productos struct {
	Nombre       string   `firestore:"nombre"`
	Precio       string   `firestore:"precio"`
	Ingredientes []string `firestore:"ingredientes"`
	Stock        string   `firestore:"stock"`
	Categoria    string   `firestore:"categoria"`
}
type ProductosResponse struct {
	ID           string   `json:"id"`
	Nombre       string   `json:"nombre"`
	Precio       string   `json:"precio"`
	Ingredientes []string `json:"ingredientes"`
	Stock        string   `json:"stock"`
	Categoria    string   `json:"categoria"`
}

// Path: entity/categorias.go

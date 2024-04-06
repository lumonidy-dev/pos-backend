package entity

import "cloud.google.com/go/firestore"

type Categorias struct {
	Nombre string                   `firestore:"nombre"`
	Desc   string                   `firestore:"desc"`
	URL    string                   `firestore:"url"`
	Tipos  []*firestore.DocumentRef `firestore:"tipos"`
}

// Path: entity/categorias.go

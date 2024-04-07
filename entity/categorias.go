package entity

// Categorias es una estructura que representa a la colección de categorías en Firestore
// @title Categorias
// @description Estructura de datos para las categorías de productos
// @type Categorias
// @prop nombre - Nombre de la categoría
// @prop desc - Descripción de la categoría
// @prop url - URL de la imagen de la categoría
// @prop tipos - Tipos de productos que pertenecen a la categoría
type Categorias struct {
	Nombre string   `firestore:"nombre"`
	Desc   string   `firestore:"desc"`
	URL    string   `firestore:"url"`
	Tipos  []string `firestore:"tipos"`
}

// CategoriasResponse es una estructura que representa a la colección de categorías en Firestore
// @title CategoriasResponse
// @description Estructura de datos para las categorías de productos
// @type CategoriasResponse
// @prop id - ID de la categoría
// @prop nombre - Nombre de la categoría
// @prop desc - Descripción de la categoría
// @prop url - URL de la imagen de la categoría
// @prop tipos - Tipos de productos que pertenecen a la categoría
type CategoriasResponse struct {
	ID     string   `json:"id"`
	Nombre string   `json:"nombre"`
	Desc   string   `json:"desc"`
	URL    string   `json:"url"`
	Tipos  []string `json:"tipos"`
}

// Path: entity/categorias.go

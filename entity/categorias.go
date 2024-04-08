package entity

// Categorias es una estructura que representa a la colección de categorías en Firestore
// @summary Estructura de datos para las categorías de productos
// @type Categorias
// @prop nombre - Nombre de la categoría
// @prop desc - Descripción de la categoría
// @prop url - URL de la imagen de la categoría
// @prop tipos - Tipos de productos que pertenecen a la categoría (nombre) y su ID
type Categorias struct {
	Nombre string   `firestore:"nombre"`
	Desc   string   `firestore:"desc"`
	URL    string   `firestore:"url"`
	Tipos  []string `firestore:"tipos"`
}

// CategoriasResponse es una estructura que representa la respuesta de la API para las categorías de productos
// @summary Estructura de datos para las categorías de productos
// @type CategoriasResponse
// @prop id - ID de la categoría
// @prop nombre - Nombre de la categoría
// @prop desc - Descripción de la categoría
// @prop url - URL de la imagen de la categoría
// @prop tipos - Tipos de productos que pertenecen a la categoría (nombre) y su ID
type CategoriasResponse struct {
	ID     string         `json:"id"`
	Nombre string         `json:"nombre"`
	Desc   string         `json:"desc"`
	URL    string         `json:"url"`
	Tipos  []TipoProducto `json:"tipos"`
}

type TipoProducto struct {
	ID           string   `json:"id"`
	Nombre       string   `json:"nombre"`
	Precio       string   `json:"precio"`
	Stock        string   `json:"stock,omitempty"`
	Ingredientes []string `json:"ingredientes"`
}

// Path: entity/categorias.go

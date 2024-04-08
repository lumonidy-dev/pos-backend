package entity

// Productos es una estructura que representa a la colección de productos en Firestore
// @summary Estructura de datos para los productos
// @type Productos
// @prop nombre - Nombre del producto
// @prop precio - Precio del producto
// @prop ingredientes - Ingredientes del producto
// @prop stock - Stock del producto puede no estar presente en la base de datos
// @prop categoria - Categoría del producto (nombre) y su ID
type Productos struct {
	Nombre       string   `firestore:"nombre"`
	Precio       string   `firestore:"precio"`
	Stock        string   `firestore:"stock" json:",omitempty"`
	Ingredientes []string `firestore:"ingredientes"`

	Categoria string `firestore:"categoria"`
}

// ProductosResponse es una estructura que representa la respuesta de la API para los productos
// @summary Estructura de datos para los productos
// @type ProductosResponse
// @prop id - ID del producto
// @prop nombre - Nombre del producto
// @prop precio - Precio del producto
// @prop ingredientes - Ingredientes del producto
// @prop stock - Stock del producto
// @prop categoria - Categoría del producto (nombre) y su ID
type ProductosResponse struct {
	ID           string    `json:"id"`
	Nombre       string    `json:"nombre"`
	Precio       string    `json:"precio"`
	Ingredientes []string  `json:"ingredientes"`
	Stock        string    `json:"stock,omitempty"`
	Categoria    Categoria `json:"categoria"`
}

type Categoria struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

// Path: entity/categorias.go

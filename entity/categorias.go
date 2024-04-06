package entity

import "cloud.google.com/go/firestore"

// Categorias representa la estructura de datos de una categoría.
type Categorias struct {
	Nombre string                   `firestore:"nombre"`
	Desc   string                   `firestore:"desc"`
	URL    string                   `firestore:"url"`
	Tipos  []*firestore.DocumentRef `firestore:"tipos"`
}

// CategoriasResponse representa la estructura de respuesta de una categoría.
type CategoriasResponse struct {
	ID     string   `json:"id"`
	Nombre string   `json:"nombre"`
	Desc   string   `json:"desc"`
	URL    string   `json:"url"`
	Tipos  []string `json:"tipos"`
}

// Definitions

/**
 * @swagger
 * definitions:
 *   Categoria:
 *     type: object
 *     properties:
 *       nombre:
 *         type: string
 *       desc:
 *         type: string
 *       url:
 *         type: string
 *       tipos:
 *         type: array
 *         items:
 *           type: string
 *
 *   CategoriaResponse:
 *     type: object
 *     properties:
 *       id:
 *         type: string
 *       nombre:
 *         type: string
 *       desc:
 *         type: string
 *       url:
 *         type: string
 *       tipos:
 *         type: array
 *         items:
 *           type: string
 */

// Path: entity/categorias.go

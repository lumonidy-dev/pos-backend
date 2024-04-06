package routes

import (
	handlers "pos-backend/handlers/categorias"
	services "pos-backend/services/categorias"

	"github.com/gorilla/mux"
)

// SetCategoriasRoutes configura las rutas para las categorías
func SetCategoriasRoutes(router *mux.Router, categoryService services.CategoryService) {
	/**
	* @swagger
	* /categorias:
	*   post:
	*     summary: Crear una categoría
	*     description: Crea una nueva categoría con la información proporcionada
	* 	parameters:
	*       - in: body
	*         name: categoria
	*         description: La categoría a crear
	*         required: true
	*         schema:
	* 		  $ref: '#/definitions/Categoria'
	*     responses:
	*       201:
	*         description: Categoría creada con éxito
	*         schema:
	*           $ref: '#/definitions/Categoria'
	*       400:
	*         description: Error al decodificar el cuerpo de la solicitud
	*       500:
	*         description: Error al crear la categoría
	 */
	router.HandleFunc("/categorias", handlers.CreateCategoryHandler(categoryService)).Methods("POST")
	/**
	 * @swagger
	 * /categorias:
	 *   get:
	 *     summary: Obtener todas las categorías
	 *     description: Obtiene una lista de todas las categorías disponibles
	 *     responses:
	 *       200:
	 *         description: Lista de todas las categorías
	 *         schema:
	 *           type: array
	 *           items:
	 *             $ref: '#/definitions/CategoriaResponse'
	 *       500:
	 *         description: Error al obtener las categorías
	 */
	router.HandleFunc("/categorias", handlers.GetAllCategoriesHandler(categoryService)).Methods("GET")
	/**
	 * @swagger
	 * /categorias/{id}:
	 *   get:
	 *     summary: Obtener una categoría por ID
	 *     description: Obtiene una categoría específica por su ID
	 *     parameters:
	 *       - in: path
	 *         name: id
	 *         description: ID de la categoría a obtener
	 *         required: true
	 *         type: string
	 *     responses:
	 *       200:
	 *         description: Categoría encontrada
	 *         schema:
	 *           $ref: '#/definitions/CategoriaResponse'
	 *       404:
	 *         description: Categoría no encontrada
	 *       500:
	 *         description: Error al obtener la categoría
	 */
	router.HandleFunc("/categorias/{id}", handlers.GetCategoryByIDHandler(categoryService)).Methods("GET")
	router.HandleFunc("/categorias/{id}", handlers.UpdateCategoryHandler(categoryService)).Methods("PUT")
	router.HandleFunc("/categorias/{id}", handlers.DeleteCategoryHandler(categoryService)).Methods("DELETE")
}

// Path: routes/categorias.go

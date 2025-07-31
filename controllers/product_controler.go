package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"texbackend/models"

	"github.com/gorilla/mux"
)

// @Summary Obtener todos los productos
// @Description Retorna una lista de productos
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetAllProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT id, nombre, descripcion, precio, imagen, stock, publicado, destacado FROM productos`)
		if err != nil {
			http.Error(w, "Error al consultar productos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var p models.Product
			err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Imagen, &p.Stock, &p.Publicar, &p.Destacado)
			if err != nil {
				http.Error(w, "Error al escanear producto", http.StatusInternalServerError)
				return
			}
			products = append(products, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

// @Summary Crear un producto
// @Description Crea un nuevo producto
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Producto a crear"
// @Success 201 {object} models.Product
// @Router /products [post]
func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO productos (nombre, descripcion, precio, imagen, stock, publicado, destacado)
				  VALUES (?, ?, ?, ?, ?, ?, ?)`
		result, err := db.Exec(query, product.Nombre, product.Descripcion, product.Precio, product.Imagen,
			product.Stock, product.Publicar, product.Destacado)
		if err != nil {
			http.Error(w, "Error al insertar producto", http.StatusInternalServerError)
			return
		}

		id, _ := result.LastInsertId()
		product.ID = int(id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

// @Summary Eliminar un producto
// @Description Elimina un producto por ID
// @Tags products
// @Param id path int true "ID del producto"
// @Success 204
// @Router /products/{id} [delete]
func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM productos WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error al eliminar producto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

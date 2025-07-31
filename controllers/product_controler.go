package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"texbackend/models"
)

// GetAllProducts godoc
// @Summary Listar productos
// @Description Obtiene todos los productos desde la base de datos
// @Tags productos
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetAllProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT 
				id, nombre, descripcion, precio, precio_descuento, imagen, 
				stock, fecha_creacion, fecha_actualizacion, publicado, destacado 
			FROM products`)
		if err != nil {
			http.Error(w, "Error al consultar productos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var productos []models.Product
		for rows.Next() {
			var p models.Product
			err := rows.Scan(
				&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.PrecioDesc,
				&p.Imagen, &p.Stock, &p.FechaCreacion, &p.FechaActual, &p.Publicar, &p.Destacado,
			)
			if err != nil {
				http.Error(w, "Error al leer productos", http.StatusInternalServerError)
				return
			}
			productos = append(productos, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productos)
	}
}

// CreateProduct godoc
// @Summary Crear producto
// @Description Inserta un nuevo producto en la base de datos
// @Tags productos
// @Accept json
// @Produce json
// @Param producto body models.Product true "Producto a crear"
// @Success 201 {object} map[string]string
// @Router /products [post]
func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		query := `
			INSERT INTO products (
				nombre, descripcion, precio, precio_descuento, imagen,
				stock, fecha_creacion, fecha_actualizacion, publicado, destacado
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		_, err := db.Exec(query,
			p.Nombre, p.Descripcion, p.Precio, p.PrecioDesc, p.Imagen,
			p.Stock, p.FechaCreacion, p.FechaActual, p.Publicar, p.Destacado,
		)
		if err != nil {
			http.Error(w, "Error al insertar producto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Producto creado con éxito"})
	}
}

// DeleteProduct godoc
// @Summary Eliminar producto
// @Description Elimina un producto por su ID
// @Tags productos
// @Param id path int true "ID del producto"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error al eliminar producto", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Producto eliminado"})
	}
}

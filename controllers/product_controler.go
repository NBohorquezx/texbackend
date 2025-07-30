package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
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
		rows, err := db.Query("SELECT id, nombre, precio FROM products")
		if err != nil {
			http.Error(w, "Error al consultar productos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var productos []models.Product
		for rows.Next() {
			var p models.Product
			if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio); err != nil {
				http.Error(w, "Error al leer productos", http.StatusInternalServerError)
				return
			}
			productos = append(productos, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productos)
	}
}

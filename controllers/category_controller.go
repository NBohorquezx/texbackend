package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"texbackend/models"

	"github.com/gorilla/mux"
)

// GetAllCategories godoc
// @Summary Listar categorías
// @Description Obtiene todas las categorías desde la base de datos
// @Tags categorias
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetAllCategories(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nombre, descripcion, activo, fecha_creacion, fecha_actualizacion FROM categories")
		if err != nil {
			http.Error(w, "Error al consultar categorías", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var categorias []models.Category
		for rows.Next() {
			var c models.Category
			err := rows.Scan(&c.ID, &c.Nombre, &c.Descripcion, &c.Activo, &c.FechaCreacion, &c.FechaActualizacion)
			if err != nil {
				http.Error(w, "Error al leer categorías", http.StatusInternalServerError)
				return
			}
			categorias = append(categorias, c)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categorias)
	}
}

// CreateCategory godoc
// @Summary Crear categoría
// @Description Crea una nueva categoría
// @Tags categorias
// @Accept json
// @Produce json
// @Param category body models.Category true "Categoría"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /categories [post]
func CreateCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO categories (nombre, descripcion, activo, fecha_creacion, fecha_actualizacion)
				  VALUES (?, ?, ?, datetime('now'), datetime('now'))`

		_, err := db.Exec(query, c.Nombre, c.Descripcion, c.Activo)
		if err != nil {
			http.Error(w, "Error al insertar categoría", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Categoría creada exitosamente"})
	}
}

// DeleteCategory godoc
// @Summary Eliminar categoría
// @Description Elimina una categoría por su ID
// @Tags categorias
// @Produce json
// @Param id path int true "ID de la categoría"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM categories WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error al eliminar la categoría", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Categoría eliminada exitosamente"})
	}
}

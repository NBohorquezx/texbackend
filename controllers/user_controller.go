package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"texbackend/models"
)

// GET: Obtener todos los usuarios
func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nombre, email FROM users")
		if err != nil {
			http.Error(w, "Error al consultar usuarios", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var u models.User
			err := rows.Scan(&u.ID, &u.Nombre, &u.Email)
			if err != nil {
				http.Error(w, "Error al leer usuario", http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// POST: Crear un nuevo usuario
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO users (nombre, email) VALUES (?, ?)`
		_, err = db.Exec(query, u.Nombre, u.Email)
		if err != nil {
			http.Error(w, "Error al insertar usuario", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario creado con éxito"})
	}
}

// DELETE: Eliminar usuario por ID
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error al eliminar usuario", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario eliminado"})
	}
}

package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID         int
	Nombre     string
	Correo     string
	Contraseña string
	Rol        string
	Activo     bool
}


func CreateUser(db *sql.DB, correo, nombre, contrasena, rol string, activo bool) error {
	query := `INSERT INTO users (correo, nombre, contrasena, rol, activo) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, correo, nombre, contrasena, rol, activo)
	if err != nil {
		return fmt.Errorf("❌ Error al insertar usuario: %v", err)
	}
	return nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, nombre, correo, contraseña, rol, activo FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Nombre, &u.Correo, &u.Contraseña, &u.Rol, &u.Activo)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	var err error
	dsn := "root:nicolas1@tcp(127.0.0.1:3306)/tienda_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %v", err)
	}

	fmt.Println("✅ Conexión a la base de datos exitosa")
	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		panic("❌ Base de datos no inicializada. Llama a InitDB() primero.")
	}
	return db
}


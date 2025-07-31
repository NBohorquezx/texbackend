package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() error {
	var err error

	// Leer el tipo de base de datos desde una variable de entorno
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite" // Valor por defecto
	}

	var dsn string

	switch dbType {
	case "mysql":
		dsn = os.Getenv("MYSQL_DSN")
		if dsn == "" {
			dsn = "root:nicolas1@tcp(127.0.0.1:3306)/tienda_db"
		}
	case "sqlite":
		dsn = "./data.db" // Usa archivo local
	default:
		return fmt.Errorf("❌ Tipo de base de datos '%s' no soportado", dbType)
	}

	// Conectar a la base de datos
	db, err = sql.Open(dbType, dsn)
	if err != nil {
		return fmt.Errorf("❌ Error al abrir la conexión: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("❌ Error al hacer ping a la base de datos: %v", err)
	}

	fmt.Printf("✅ Conectado a la base de datos usando %s\n", dbType)
	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		panic("❌ Base de datos no inicializada. Llama a InitDB() primero.")
	}
	return db
}

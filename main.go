package main

import (
	"log"
	"net/http"

	"texbackend/config"
	"texbackend/routes"

	_ "texbackend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API Tienda Online
// @version 1.0
// @description Esta es la API de productos para la tienda.
// @host localhost:8080
// @BasePath /
func main() {
	// Inicializa la conexión
	if err := config.InitDB(); err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	// Obtiene la instancia *sql.DB
	db := config.GetDB()

	// Define rutas y servidor
	router := routes.SetupRoutes(db)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

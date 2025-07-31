package main

import (
	"log"
	"net/http"
	"os"

	"texbackend/config"
	"texbackend/routes"

	_ "texbackend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API Tienda Online
// @version 1.0
// @description Esta es la API de productos para la tienda.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Inicializa base de datos (sqlite por defecto)
	if err := config.InitDB(); err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}
	db := config.GetDB()

	// Cargar rutas
	router := routes.SetupRoutes(db)

	// Documentación Swagger en /swagger/index.html
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor iniciado en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

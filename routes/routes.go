package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"texbackend/controllers"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Ruta base
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🛒 Bienvenido a la API de la tienda"))
	})

	// Rutas de productos
	router.HandleFunc("/products", controllers.GetAllProducts(db)).Methods("GET")

	return router
}

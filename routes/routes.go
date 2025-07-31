package routes

import (
	"database/sql"
	"net/http"

	"texbackend/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes define todas las rutas de la API
func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Prefijo de versión si lo deseas (puedes comentar si no lo usas)
	api := router.PathPrefix("/api/v1").Subrouter()

	// Productos
	api.HandleFunc("/products", controllers.GetAllProducts(db)).Methods("GET")
	api.HandleFunc("/products", controllers.CreateProduct(db)).Methods("POST")
	api.HandleFunc("/products/{id:[0-9]+}", controllers.DeleteProduct(db)).Methods("DELETE")

	// Usuarios (si tienes)
	api.HandleFunc("/users", controllers.GetAllUsers(db)).Methods("GET")
	api.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser(db)).Methods("DELETE")

	// Categorías (si tienes)
	api.HandleFunc("/categories", controllers.GetAllCategories(db)).Methods("GET")
	api.HandleFunc("/categories", controllers.CreateCategory(db)).Methods("POST")
	api.HandleFunc("/categories/{id:[0-9]+}", controllers.DeleteCategory(db)).Methods("DELETE")

	// Endpoint de salud
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")

	return router
}

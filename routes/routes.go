package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"texbackend/controllers"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	//Productos
	router.HandleFunc("/products", controllers.GetAllProducts(db)).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct(db)).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.DeleteProduct(db)).Methods("DELETE")

	//Usuarios
	router.HandleFunc("/users", controllers.GetAllUsers(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser(db)).Methods("DELETE")

	//Categorías
	router.HandleFunc("/categories", controllers.GetAllCategories(db)).Methods("GET")
	router.HandleFunc("/categories", controllers.CreateCategory(db)).Methods("POST")
	router.HandleFunc("/categories/{id:[0-9]+}", controllers.DeleteCategory(db)).Methods("DELETE")

	// 🩺 Salud del servidor
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	return router
}
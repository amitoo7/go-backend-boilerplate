package routes

import (
	"backend-boilerplate/handlers"
	"backend-boilerplate/middleware"

	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You are authorized!"))
	}).Methods("GET")

	return router
}

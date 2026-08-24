package api

import (
	"goapi/api/handlers"
	"goapi/api/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configure les routes de l'API
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	// Routes publiques (pas besoin d'authentification)
	r.HandleFunc("/api/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/delete-users", handlers.DeleteUsers).Methods("POST")

	// Routes protégées (nécessitent un token JWT valide)
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/api/logout", handlers.Logout).Methods("POST")
	protected.HandleFunc("/api/profile", handlers.GetProfile).Methods("GET")

	return r
}

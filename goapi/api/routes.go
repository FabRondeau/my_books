package api

import (
	"fmt"

	"github.com/gorilla/mux"
)

// SetupRoutes configure les routes de l'API
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	// Ajoute d'autres routes ici
	fmt.Println("SetupRoutes")
	return r
}

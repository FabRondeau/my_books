package main

import (
	"goapi/db"
	"log"
	"net/http"
)

func main() {
	// Initialise la base de données
	if err := db.InitDB(); err != nil {
		log.Fatal("Erreur lors de l'initialisation de la DB :", err)
	}
	defer db.CloseDB()

	// Configure les routes de l'APImon
	// Démarre le serveur
	log.Println("Serveur démarré sur :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

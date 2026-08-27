package main

import (
	"goapi/config"
	"goapi/controller"
	"goapi/db"
	"goapi/router"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Initialisation de la base de données
	db.InitDB()
	defer db.CloseDB()
	dbs := config.DatabaseConnection()

	validate := validator.New()

	// Initialisation des contrôleurs
	authController := controller.NewAuthControllerImpl(dbs, validate)
	publicController := controller.NewPublicControllerImpl(dbs, validate)
	publisherController := controller.NewPublisherControllerImpl(dbs, validate)
	authorController := controller.NewAuthorControllerImpl(dbs, validate)

	// Création d'une instance unique de gin.Engine
	engine := gin.Default()

	// Ajout des routes à l'instance principale
	router.AuthRouter(engine, authController)
	router.PublicRouter(engine, publicController)
	router.PublisherRouter(engine, publisherController)
	router.AuthorRouter(engine, authorController)

	// Configuration du serveur HTTP
	server := &http.Server{
		Addr:           ":" + os.Getenv("API_PORT"),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Démarrage du serveur
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

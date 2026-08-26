package main

import (
	"goapi/config"
	"goapi/controller"
	"goapi/db"
	"goapi/router"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Database
	db.InitDB()
	db.CloseDB()
	db := config.DatabaseConnection()

	validate := validator.New()

	// Controller
	authController := controller.NewAuthControllerImpl(db, validate)

	// Router
	routes := router.AuthRouter(authController)

	server := &http.Server{
		Addr:           ":" + os.Getenv("API_PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

package db

import (
	"goapi/config"
	"goapi/model"

	_ "github.com/mattn/go-sqlite3"
)

// DB est une instance globale de la base de données
var DB DBInterface

// InitDB initialise la connexion à la base de données
func InitDB() error {
	conn := config.DatabaseConnection()
	conn.AutoMigrate(model.User{})
	conn.AutoMigrate(model.Publisher{})
	conn.AutoMigrate(model.Author{})
	conn.AutoMigrate(model.Book{})

	return nil
}

// CloseDB ferme la connexion et réinitialise DB à nil
func CloseDB() {
	if DB != nil {
		DB.Close()
		DB = nil // Réinitialise DB à nil après fermeture
	}
}

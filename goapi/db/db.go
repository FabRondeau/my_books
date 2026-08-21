package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Driver SQLite
)

// DB est une instance globale de la base de données
var DB *sql.DB

// InitDB initialise la connexion à la base de données
func InitDB() error {
	var err error
	db_name := "library"
	DB, err = sql.Open("sqlite3", "./"+db_name+".sqlite")
	if err != nil {
		return err
	}
	// Vérifie la connexion
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}

// CloseDB ferme la connexion
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

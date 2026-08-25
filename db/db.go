package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB est une instance globale de la base de données
var DB DBInterface

// InitDB initialise la connexion à la base de données
func InitDB() error {
	var dbname = "library"
	conn, err := sql.Open("sqlite3", "./"+dbname+".sqlite")
	if err != nil {
		return err
	}

	// Vérifie la connexion
	if err = conn.Ping(); err != nil {
		return err
	}

	// Exécute les migrations
	if err = RunMigrations(conn); err != nil {
		return err
	}
	// Assigne l'interface
	DB = conn
	return nil
}

// CloseDB ferme la connexion et réinitialise DB à nil
func CloseDB() {
	if DB != nil {
		DB.Close()
		DB = nil // Réinitialise DB à nil après fermeture
	}
}

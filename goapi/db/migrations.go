package db

import (
	"database/sql"
)

// RunMigrations exécute les migrations nécessaires pour la base de données
func RunMigrations(db *sql.DB) error {
	// Crée la table `users`
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        );
    `)
	if err != nil {
		return err
	}

	// Ajoute d'autres migrations ici si nécessaire
	return nil
}

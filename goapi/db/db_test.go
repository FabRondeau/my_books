package db

import (
	"database/sql"
	"testing"
)

func TestCloseDB(t *testing.T) {
	// Sauvegarde l'instance actuelle de DB
	oldDB := DB

	// Crée une connexion factice pour le test
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Échec de la création de la DB en mémoire : %v", err)
	}

	// Teste la fermeture
	CloseDB()

	// Vérifie que DB est nil après fermeture
	if DB != nil {
		t.Errorf("DB n'est pas nil après fermeture, valeur actuelle : %v", DB)
	}

	// Restaure l'ancienne valeur
	DB = oldDB
}

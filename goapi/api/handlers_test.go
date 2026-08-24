package api

import (
	"encoding/json"
	"goapi/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUsersWithMock(t *testing.T) {
	// Crée une mock de la base de données avec sqlmock
	mockDB, mock, err := db.NewMockDB()
	if err != nil {
		t.Fatalf("Échec de la création de la mock DB : %v", err)
	}
	defer mockDB.Close()

	// Sauvegarde l'instance actuelle de DB
	oldDB := db.DB
	db.DB = mockDB                   // Assigne la mock à la variable globale DB
	defer func() { db.DB = oldDB }() // Restaure l'ancienne valeur après le test

	// Configure la mock pour retourner des utilisateurs simulés
	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(1, "testuser", "test@example.com").
		AddRow(2, "anotheruser", "another@example.com")

	mock.ExpectQuery("SELECT id, username, email FROM user").
		WillReturnRows(rows)

	// Crée une requête HTTP
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crée un ResponseRecorder pour capturer la réponse
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)

	// Appelle le handler
	handler.ServeHTTP(rr, req)

	// Vérifie le code de statut
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Code de statut attendu : %v, obtenu : %v", http.StatusOK, status)
	}

	// Vérifie le corps de la réponse
	var users []db.User
	err = json.NewDecoder(rr.Body).Decode(&users)
	if err != nil {
		t.Fatalf("Échec du décodage de la réponse : %v", err)
	}

	// Vérifie que les utilisateurs simulés sont retournés
	if len(users) != 2 {
		t.Errorf("Nombre d'utilisateurs attendu : 2, obtenu : %d", len(users))
	}

	// Vérifie que toutes les attentes de la mock ont été satisfaites
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Attentes de la mock non satisfaites : %v", err)
	}
}

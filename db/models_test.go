package db

import (
	"testing"
)

func TestUserStruct(t *testing.T) {
	user := User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	// Vérifie que les champs sont correctement assignés
	if user.ID != 1 {
		t.Errorf("ID attendu : 1, obtenu : %d", user.ID)
	}
	if user.Username != "testuser" {
		t.Errorf("Username attendu : testuser, obtenu : %s", user.Username)
	}
	if user.Email != "test@example.com" {
		t.Errorf("Email attendu : test@example.com, obtenu : %s", user.Email)
	}
}

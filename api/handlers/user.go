package handlers

import (
	"encoding/json"
	"goapi/api/middleware"
	"goapi/db"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID utilisateur depuis le contexte avec UserIDKey
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var user db.User
	err := db.DB.QueryRow(
		"SELECT id, username, email FROM user WHERE id = ?",
		userID,
	).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Utilisateur non trouvé"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"user":    user,
	})
}

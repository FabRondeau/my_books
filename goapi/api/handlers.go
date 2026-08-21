package api

import (
	"encoding/json"
	"goapi/db"
	"net/http"
)

// GetUsers retourne la liste des utilisateurs
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []db.User
	for rows.Next() {
		var u db.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

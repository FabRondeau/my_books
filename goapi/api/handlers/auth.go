package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/utils"
	"net/http"
)

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

// Fix empty password when Signin up (because of password string json: "-")
type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // Autorise la désérialisation
}

// SignUp gère l'inscription d'un nouvel utilisateur
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Fix empty password when Signin up (because of password string json: "-")
	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"success": false, "message": "Données invalides"}`, http.StatusBadRequest)
		return
	}
	// Fix empty password when Signin up (because of password string json: "-")
	// Crée un utilisateur avec les données de la requête
	user := db.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // Copie le mot de passe pour le hacher
	}

	// Hache le mot de passe
	if err := user.HashPassword(); err != nil {
		http.Error(w, `{"success": false, "message": "Erreur lors du hachage du mot de passe"}`, http.StatusInternalServerError)
		return
	}

	// Insère l'utilisateur dans la base de données
	_, err := db.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, user.Password,
	)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Nom d'utilisateur ou email déjà utilisé"}`, http.StatusConflict)
		return
	}

	// Génère un token JWT
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Erreur lors de la génération du token"}`, http.StatusInternalServerError)
		return
	}

	// Retourne le token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{
		Success: true,
		Message: "Inscription réussie",
		Token:   token,
	})
}

// Login gère la connexion d'un utilisateur
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, `{"success": false, "message": "Données invalides"}`, http.StatusBadRequest)
		return
	}

	// Récupère l'utilisateur depuis la base de données
	var user db.User
	err := db.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", credentials.Username).
		Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Utilisateur non trouvé."}`, http.StatusUnauthorized)
		return
	}

	// Vérifie le mot de passe
	if err := user.CheckPassword(credentials.Password); err != nil {
		// http.Error(w, `{`+err.Error()+`}`, http.StatusUnauthorized)
		http.Error(w, `{"success": false, "message": "Nom d'utilisateur ou mot de passe incorrect"}`, http.StatusUnauthorized)
		return
	}

	// Génère un token JWT
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Erreur lors de la génération du token"}`, http.StatusInternalServerError)
		return
	}

	// Retourne le token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{
		Success: true,
		Message: "Connexion réussie",
		Token:   token,
	})
}

// Logout gère la déconnexion (invalide le token côté client)
func Logout(w http.ResponseWriter, r *http.Request) {
	// En JWT, la déconnexion est gérée côté client (suppression du token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{
		Success: true,
		Message: "Déconnexion réussie",
	})
}

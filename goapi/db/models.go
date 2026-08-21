package db

// User représente un utilisateur dans la base de données
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

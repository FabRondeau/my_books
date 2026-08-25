package db

import (
	"golang.org/x/crypto/bcrypt"
)

// User représente un utilisateur dans la base de données
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // Désérialise mais n'inclut pas dans les réponses si vide
}

type Book struct {
	ID            int    `json:"id"`
	ISBN13        string `json:"isbn13"`
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Kind          string `json:"kind"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"publisheddate"`
}

type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HashPassword hache le mot de passe avec bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword vérifie si le mot de passe fourni correspond au hash
// TODO - mots de passe
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

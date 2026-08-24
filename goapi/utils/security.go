package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("ta_clé_secrète_à_changer") // À stocker dans des variables d'environnement

// GenerateJWT génère un token JWT pour un utilisateur
func GenerateJWT(userID int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Expiration : 24h
	})

	return token.SignedString(jwtSecret)
}

// ValidateJWT valide un token JWT et retourne les claims
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

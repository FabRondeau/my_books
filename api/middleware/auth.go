package middleware

import (
	"context"
	"goapi/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

// Exporte la constante en la renommant avec une majuscule
const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"success": false, "message": "Token manquant"}`, http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, `{"success": false, "message": "Format du token invalide"}`, http.StatusUnauthorized)
			return
		}

		token, err := utils.ValidateJWT(parts[1])
		if err != nil || !token.Valid {
			http.Error(w, `{"success": false, "message": "Token invalide ou expiré"}`, http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		// Utilise UserIDKey au lieu de userIDKey
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

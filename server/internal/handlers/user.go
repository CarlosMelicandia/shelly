package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(config.LoadEnv().JWTSecret)
		return secret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract the userId claim
	var userId string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if uid, ok := claims["userId"].(string); ok {
			userId = uid
		} else {
			http.Error(w, "userId not found in token claims", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	// todo: remove this and only send the users info for the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":  tokenString,
		"userId": userId,
	})
}

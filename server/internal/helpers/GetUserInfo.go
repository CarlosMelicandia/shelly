package helpers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
)

func GetToken(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, fmt.Errorf("missing token: %w", err)
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
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return token, nil
}

func GetTokenString(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", fmt.Errorf("missing token: %w", err)
	}
	return cookie.Value, nil
}

func GetUserId(r *http.Request) (string, error) {
	token, err := GetToken(r)
	if err != nil {
		return "", fmt.Errorf("could not get the token to retrieve userId: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(string); ok {
			return userId, nil
		}
	}

	return "", fmt.Errorf("could not get the UserId")
}

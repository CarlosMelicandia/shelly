package token

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
)

func parseToken(tokenString string, tokenType TokenType) (*jwt.Token, error) {
	var secret []byte

	// Determine the secret based on the token type
	switch tokenType {
	case AccessToken:
		secret = []byte(config.LoadEnv().JWTSecret)
	case RefreshToken:
		secret = []byte(config.LoadEnv().JWTSecretRefresh)
	default:
		return nil, fmt.Errorf("invalid token type")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token parsing error: %w", err)
	}

	// Check for expiration
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if time.Now().After(expirationTime) {
				return nil, fmt.Errorf("token is expired")
			}
		} else {
			return nil, fmt.Errorf("token does not contain a valid 'exp' claim")
		}
	}

	return token, nil
}

func refreshTokensAndParseAccessToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	newAccessToken, err := refreshTokens(w, r)
	if err != nil {
		return nil, fmt.Errorf("refresh failed: %w", err)
	}

	token, err := parseToken(newAccessToken, AccessToken)
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid access token after refresh: %w", err)
	}

	return token, nil
}

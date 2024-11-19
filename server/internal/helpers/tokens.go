package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	oauth "github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
)


func ValidateAccessToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", fmt.Errorf("missing access token")
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.LoadEnv().JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid or expired token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if time.Now().After(expirationTime) {
				return "", fmt.Errorf("token has expired")
			}
		} else {
			return "", fmt.Errorf("token is missing expiration claim")
		}
	}

	return tokenString, nil
}

func RefreshTokens(w http.ResponseWriter, r *http.Request) (string, error) {
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		return "", fmt.Errorf("missing access and/or refresh token")
	}

	refreshTokenString := refreshCookie.Value
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.LoadEnv().JWTSecretRefresh), nil
	})

	if err != nil || !refreshToken.Valid {
		return "", fmt.Errorf("invalid or expired refresh token, please log in again")
	}

	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if time.Now().After(expirationTime) {
				return "", fmt.Errorf("refresh token has expired, please log in again")
			}
		} else {
			return "", fmt.Errorf("refresh token is missing expiration claim")
		}

		userId, ok := claims["userId"].(string)
		if !ok {
			return "", fmt.Errorf("invalid refresh token claims")
		}

		newAccessToken, err := oauth.GenerateJWT(userId)
		if err != nil {
			return "", fmt.Errorf("failed to generate new access token")
		}

		newRefreshToken, err := oauth.GenerateRefreshToken(userId)
		if err != nil {
			return "", fmt.Errorf("failed to generate new refresh token")
		}

		oauth.SetCookie(w, "access_token", newAccessToken, time.Now().Add(15*time.Minute))
		oauth.SetCookie(w, "refresh_token", newRefreshToken, time.Now().AddDate(0, 3, 0))

		return newAccessToken, nil
	}

	return "", fmt.Errorf("invalid refresh token claims")
}

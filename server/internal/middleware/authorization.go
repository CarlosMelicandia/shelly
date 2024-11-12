package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/auth"
	"github.com/weareinit/Opal/internal/config"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := validateAccessToken(r)
		if err != nil {
			// Attempt to refresh the access token using the refresh token
			newAccessToken, err := refreshTokens(w, r)
			if err != nil {
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "access_token", newAccessToken)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx := context.WithValue(r.Context(), "access_token", accessToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateAccessToken(r *http.Request) (string, error) {
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

func refreshTokens(w http.ResponseWriter, r *http.Request) (string, error) {
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

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from your frontend origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4321")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

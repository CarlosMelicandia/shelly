package tokens

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/auth/utils"
	"github.com/weareinit/Opal/internal/config"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token, err := ValidateAccessToken(w, r)
	if err != nil {
		newAccessToken, err := refreshTokensAndParseAccessToken(w, r)
		if err != nil {
			return nil, fmt.Errorf("could not refresh the access token: %w", err)
		}
		return newAccessToken, nil
	}
	return token, nil
}

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
	newAccessToken, err := RefreshTokens(w, r)
	if err != nil {
		return nil, fmt.Errorf("refresh failed: %w", err)
	}

	token, err := parseToken(newAccessToken, AccessToken)
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid access token after refresh: %w", err)
	}

	return token, nil
}

func GetAccessTokenString(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		newToken, refreshErr := RefreshTokens(w, r)
		if refreshErr != nil {
			return "", fmt.Errorf("missing token and refresh failed: %w", refreshErr)
		}
		return newToken, nil
	}

	return cookie.Value, nil
}

func ValidateAccessToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString, err := GetAccessTokenString(w, r)
	if err != nil {
		return nil, fmt.Errorf("could not fetch token: %w", err)
	}

	token, err := parseToken(tokenString, AccessToken)
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token: %w", err)
	}

	return token, nil
}

func RefreshTokens(w http.ResponseWriter, r *http.Request) (string, error) {
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		return "", fmt.Errorf("missing access and/or refresh token")
	}

	refreshTokenString := refreshCookie.Value
	refreshToken, err := parseToken(refreshTokenString, RefreshToken)
	if err != nil || !refreshToken.Valid {
		return "", fmt.Errorf("invalid or expired refresh token, please log in again: %w", err)
	}

	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		userId, ok := claims["userId"].(string)
		if !ok {
			return "", fmt.Errorf("invalid refresh token claims")
		}

		newAccessToken, err := utils.GenerateJWT(userId)
		if err != nil {
			return "", fmt.Errorf("failed to generate new access token: %w", err)
		}

		newRefreshToken, err := utils.GenerateRefreshToken(userId)
		if err != nil {
			return "", fmt.Errorf("failed to generate new refresh token: %w", err)
		}

		utils.SetCookie(w, "access_token", newAccessToken, time.Now().Add(15*time.Minute))
		utils.SetCookie(w, "refresh_token", newRefreshToken, time.Now().AddDate(0, 3, 0))

		return newAccessToken, nil
	}

	return "", fmt.Errorf("invalid refresh token claims")
}

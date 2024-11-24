package token

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/auth/utils"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token, err := validateAccessToken(w, r)
	if err != nil {
		newAccessToken, err := refreshTokensAndParseAccessToken(w, r)
		if err != nil {
			return nil, fmt.Errorf("could not refresh the access token: %w", err)
		}
		return newAccessToken, nil
	}
	return token, nil
}

func GetAccessTokenString(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		newToken, refreshErr := refreshTokens(w, r)
		if refreshErr != nil {
			return "", fmt.Errorf("missing token and refresh failed: %w", refreshErr)
		}
		return newToken, nil
	}

	return cookie.Value, nil
}

func validateAccessToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
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

func refreshTokens(w http.ResponseWriter, r *http.Request) (string, error) {
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

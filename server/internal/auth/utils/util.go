// Pretty straight-forward: anything authentication-related util functions should be saved here.

package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
)

// Generates a JWT with a 15 min expiration time. This is good practice to have a short living access token.
// Access tokens and JWT's are words that can be used interchangeably.
func GenerateJWT(userId string) (string, error) {
	envConfig := config.LoadEnv()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
		"userId": userId,
	})

	secret := []byte(envConfig.JWTSecret)
	return token.SignedString(secret)
}

// Generates a long-living refresh token with a 3 month expiration time.
// We will use this refresh token to create a new access token (both of which are cookies that are sent to the client)
func GenerateRefreshToken(userId string) (string, error) {
	envConfig := config.LoadEnv()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().AddDate(0, 3, 0).Unix(),
		"userId": userId,
	})

	refresh_secret := []byte(envConfig.JWTSecretRefresh)
	return token.SignedString(refresh_secret)
}

// This function is used solely to help users who have been logged out to have an indication of what email they used to log in.
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}
	localPart := parts[0]
	domainPart := parts[1]

	// Mask the local part except for the first two and last two characters
	if len(localPart) > 4 {
		masked := fmt.Sprintf("%s*****%s", localPart[:2], localPart[len(localPart)-2:])
		return fmt.Sprintf("%s@%s", masked, domainPart)
	}

	// If the local part is too short to mask, return it as is
	return email
}

// Nice helper method used create a cookie on the client
func SetCookie(w http.ResponseWriter, name string, value string, expires time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		Path:     "/",
		HttpOnly: true,
	})
}

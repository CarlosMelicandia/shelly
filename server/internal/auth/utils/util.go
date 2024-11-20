package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
)

// generates a JWT with a 15 min expiration time
func GenerateJWT(userId string) (string, error) {
	envConfig := config.LoadEnv()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
		"userId": userId,
	})

	secret := []byte(envConfig.JWTSecret)
	return token.SignedString(secret)
}

// generates a long-living refresh token with a 3 month expiration time.
func GenerateRefreshToken(userId string) (string, error) {
	envConfig := config.LoadEnv()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().AddDate(0, 3, 0).Unix(),
		"userId": userId,
	})

	refresh_secret := []byte(envConfig.JWTSecretRefresh)
	return token.SignedString(refresh_secret)
}

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

func SetCookie(w http.ResponseWriter, name string, value string, expires time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		Path:     "/",
		HttpOnly: true,
	})
}

package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	GivenName     string `json:"given_name"`
	// FamilyName is the only field that can be optional
	FamilyName    *string `json:"family_name,omitempty"`
	ProfilePicURL string  `json:"picture"`
}

func InitOAuthConfig() {
	envConfig := config.LoadEnv()

	googleOauthConfig = &oauth2.Config{
		ClientID:     envConfig.GoogleClientID,
		ClientSecret: envConfig.GoogleClientSecret,
		RedirectURL:  "http://localhost:8000/api/auth/callback/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

var oauthStateString = randomState()

// randomState generates a random string to prevent CSRF attacks
func randomState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// HandleGoogleLogin redirects the user to Google's OAuth2 consent page
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	InitOAuthConfig()
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGoogleCallback handles the Google OAuth callback and returns a JWT
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		log.Println("Invalid OAuth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code exchange failed:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Fetch user info using the access token
	client := googleOauthConfig.Client(context.Background(), token)
	userInfo, err := getUserInfo(client)

	// Generate JWT after successful authentication
	jwtToken, err := GenerateJWT(userInfo.ID)
	if err != nil {
		log.Println("Error generating JWT:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	refreshToken, err := GenerateRefreshToken(userInfo.ID)
	if err != nil {
		log.Println("Error generating secret:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	SetCookie(w, "access_token", jwtToken, time.Now().Add(15*time.Minute))
	SetCookie(w, "refresh_token", refreshToken, time.Now().AddDate(0, 3, 0))

	censoredEmail := maskEmail(userInfo.Email)

	// Mask email so user can know which account they used to log in originally
	SetCookie(w, "mask_email", censoredEmail, time.Now().AddDate(0, 6, 0))

	http.Redirect(w, r, "http://localhost:8000/dashboard/", http.StatusSeeOther)
}

func getUserInfo(client *http.Client) (*GoogleUserInfo, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body
	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	// printUserInfo(userInfo)

	return &userInfo, nil
}

func printUserInfo(userInfo GoogleUserInfo) {
	userInfoJson, err := json.MarshalIndent(userInfo, "", "  ")
	if err != nil {
		fmt.Println("Error formatting user info:", err)
		return
	}
	fmt.Println("UserInfo:\n", string(userInfoJson))
}

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

func maskEmail(email string) string {
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

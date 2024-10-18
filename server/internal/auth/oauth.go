package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func InitOAuthConfig() {
	envConfig := config.LoadEnv()

	googleOauthConfig = &oauth2.Config{
		ClientID:     envConfig.GoogleClientID,
		ClientSecret: envConfig.GoogleClientSecret,
		RedirectURL:  "http://localhost:8000/api/auth/callback/google",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
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
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Error getting user info:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer userInfo.Body.Close()

	// Generate JWT after successful authentication
	jwtToken, err := generateJWT()
	if err != nil {
		log.Println("Error generating JWT:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Send the JWT to the client (usually as a cookie or JSON response)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtToken,
		Expires: time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "http://localhost:4321/dashboard", http.StatusSeeOther)
}

// generates a JWT with a 24-hour expiration time
func generateJWT() (string, error) {
	envConfig := config.LoadEnv()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := []byte(envConfig.JWTSecret)
	return token.SignedString(secret)
}

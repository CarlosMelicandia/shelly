package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/cmd/operations"
	"github.com/weareinit/Opal/internal/auth/utils"
	"github.com/weareinit/Opal/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

type GoogleUserInfo struct {
	Id            string `json:"id"`
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
	jwtToken, err := utils.GenerateJWT(userInfo.Id)
	if err != nil {
		log.Println("Error generating JWT:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(userInfo.Id)
	if err != nil {
		log.Println("Error generating secret:", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

  user := api.User {
    UserId: userInfo.Id,
    FirstName: userInfo.GivenName,
    LastName: userInfo.FamilyName,
    Email: userInfo.Email,
  }

  if _, err := operations.CreateUser(user); err != nil {
    fmt.Println("Error creating user:", err)
    return
  }

	utils.SetCookie(w, "access_token", jwtToken, time.Now().Add(15*time.Minute))
	utils.SetCookie(w, "refresh_token", refreshToken, time.Now().AddDate(0, 3, 0))

	censoredEmail := utils.MaskEmail(userInfo.Email)

	// Mask email so user can know which account they used to log in originally
	utils.SetCookie(w, "mask_email", censoredEmail, time.Now().AddDate(0, 6, 0))

	http.Redirect(w, r, "http://localhost:8000/?register_form=true", http.StatusSeeOther)
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


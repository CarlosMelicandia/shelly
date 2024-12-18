// This file is primarily completed. There are some minor additions that can be made like adding and checking for certain
// permissions depending on what roles the user has.

// For example, as of writing this (12/15/24), we are checking if the user who is logging in has they
// EBoard discord role. If so, we allow that user to have permissions to go to the /admin route. We can add
// more checks like seeing if a user has a mentor, volunteer, or sponsor role. Other than that,
// unless you know what you are doing, you probably want to leave this file alone.

package oauth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/weareinit/Opal/cmd/operations"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/internal/helpers/user"
	"github.com/weareinit/Opal/internal/tools"
)

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

type GuildMember struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
	Roles []string `json:"roles"`
}

const (
	discordAuthorizeURL        = "https://discord.com/api/oauth2/authorize"
	discordTokenURL            = "https://discord.com/api/oauth2/token"
	discordUserURL             = "https://discord.com/api/users/@me"
	discordGuildMemberEndpoint = "https://discord.com/api/v10/guilds/%s/members/%s"
)

const INIT_DISCORD = "245393533391863808"
const INIT_EBOARD_ROLE = "399558426511802368"

func HandleDiscordLogin(w http.ResponseWriter, r *http.Request) {
	envConfig := config.LoadEnv()
	url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=identify guilds.join guilds.members.read", discordAuthorizeURL, envConfig.DiscordClientID, envConfig.DiscordRedirectURI)
	http.Redirect(w, r, url, http.StatusFound)
}

func HandleDiscordCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	discordUser, err := fetchDiscordUser(token)
	if err != nil {
		http.Error(w, "Failed to fetch Discord user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := user.GetUserId(w, r)
	if err != nil {
		http.Error(w, "Failed to fetch User ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := operations.ConnectDiscordToUser(discordUser.ID, userId); err != nil {
		http.Error(w, "Failed to save Discord ID", http.StatusInternalServerError)
		return
	}

	if _, err := checkAndUpdateAdminRole(token, INIT_DISCORD, INIT_EBOARD_ROLE, userId); err != nil {
		http.Error(w, "Failed to check/update admin role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:8000/dashboard", http.StatusSeeOther)
}

func exchangeCodeForToken(code string) (string, error) {
	envConfig := config.LoadEnv()
	data := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s", envConfig.DiscordClientID, envConfig.DiscordClientSecretID, code, envConfig.DiscordRedirectURI)
	req, err := http.NewRequest("POST", discordTokenURL, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func fetchDiscordUser(accessToken string) (*DiscordUser, error) {
	req, err := http.NewRequest("GET", discordUserURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user DiscordUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func checkAndUpdateAdminRole(discordToken string, INITServer string, EBoardRole string, userId string) (any, error) {
	url := fmt.Sprintf("https://discord.com/api/users/@me/guilds/%s/member", INITServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+discordToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			// User is not on the server, refetch user information as fallback
			userInfoURL := "https://discord.com/api/users/@me"
			userReq, err := http.NewRequest("GET", userInfoURL, nil)
			if err != nil {
				return nil, err
			}
			userReq.Header.Set("Authorization", "Bearer "+discordToken)
			userReq.Header.Set("Content-Type", "application/json")

			userResp, err := client.Do(userReq)
			if err != nil {
				return nil, err
			}
			defer userResp.Body.Close()

			if userResp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(userResp.Body)
				return nil, fmt.Errorf("error fetching user info, status: %d, response: %s", userResp.StatusCode, string(body))
			}

			var discordUser DiscordUser
			err = json.NewDecoder(userResp.Body).Decode(&discordUser)
			if err != nil {
				return nil, err
			}

			// Call to revoke admin status if user is not on the server
			return checkAndRevokeAdminRole(discordToken, INITServer, EBoardRole, userId)
		}

		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch guild member data, status: %d, response: %s", resp.StatusCode, string(body))
	}

	var member GuildMember
	err = json.NewDecoder(resp.Body).Decode(&member)
	if err != nil {
		return nil, err
	}

	// Check if the user has the required role
	hasRole := false
	for _, role := range member.Roles {
		if role == EBoardRole {
			hasRole = true
			break
		}
	}

	if hasRole {
		// Update the is_admin field to 1 if the user has the required role
		return tools.LoadDB(func(db *sql.DB) (any, error) {
			updateQuery := `UPDATE user SET is_admin = 1 WHERE user_id = ?`
			_, err := db.Exec(updateQuery, userId)
			if err != nil {
				return nil, err
			}
			return "User's is_admin status successfully updated", nil
		})
	}

	return checkAndRevokeAdminRole(discordToken, INITServer, EBoardRole, userId)
}

func checkAndRevokeAdminRole(discordToken string, INITServer string, EBoardRole string, userId string) (any, error) {
	url := fmt.Sprintf("https://discord.com/api/users/@me/guilds/%s/member", INITServer)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+discordToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return tools.LoadDB(func(db *sql.DB) (any, error) {
				updateQuery := `UPDATE user SET is_admin = 0 WHERE user_id = ?`
				_, err := db.Exec(updateQuery, userId)
				if err != nil {
					return nil, err
				}
				return "User's is_admin status successfully revoked as they are not on the server", nil
			})
		}

		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch guild member data, status: %d, response: %s", resp.StatusCode, string(body))
	}

	var member GuildMember
	err = json.NewDecoder(resp.Body).Decode(&member)
	if err != nil {
		return nil, err
	}

	hasRole := false
	for _, role := range member.Roles {
		if role == EBoardRole {
			hasRole = true
			break
		}
	}

	if !hasRole {
		return tools.LoadDB(func(db *sql.DB) (any, error) {
			updateQuery := `UPDATE user SET is_admin = 0 WHERE user_id = ?`
			_, err := db.Exec(updateQuery, userId)
			if err != nil {
				return nil, err
			}
			return "User's is_admin status successfully revoked", nil
		})
	}

	return "User still has the required role, no action taken", nil
}

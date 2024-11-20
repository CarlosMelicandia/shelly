package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/helpers/tokens"
	"github.com/weareinit/Opal/internal/tools"
)


func GetUserId(w http.ResponseWriter, r *http.Request) (string, error) {
	token, err := tokens.GetAccessToken(w, r)
	if err != nil {
		return "", fmt.Errorf("could not get the token to retrieve userId: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(string); ok {
			return userId, nil
		}
	}

	return "", fmt.Errorf("could not get the UserId")
}

func GetUser(w http.ResponseWriter, r *http.Request) (api.User, error) {
    userId, err := GetUserId(w, r)
    if err != nil {
        return api.User{}, fmt.Errorf("could not get the token from user")
    }

    return tools.LoadDB(func(db *sql.DB) (api.User, error) {
        var user api.User
        getUserQuery := `SELECT user_id, name, family_name, email, discord_id, hacker_id, is_admin, is_volunteer, is_mentor, is_sponsor FROM user WHERE user_id = ?`

        row := db.QueryRow(getUserQuery, userId)

        err := row.Scan(
            &user.UserId,
            &user.Name,
            &user.FamilyName,
            &user.Email,
            &user.DiscordId,
            &user.HackerId,
            &user.IsAdmin,
            &user.IsVolunteer,
            &user.IsMentor,
            &user.IsSponsor,
        )
        if err != nil {
            if err == sql.ErrNoRows {
                return api.User{}, fmt.Errorf("no user found with the given ID")
            }
            return api.User{}, fmt.Errorf("failed to retrieve the user: %v", err)
        }

        return user, nil
    })
}

func IsUserAdmin(userId string, w http.ResponseWriter, r *http.Request) bool {
  getUser, err := GetUser(w, r)

  if err != nil {
    return false
  }

  if getUser.IsAdmin {
    return true
  } else {
    return false
  }
}

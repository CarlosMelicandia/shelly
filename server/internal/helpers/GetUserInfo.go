package helpers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/config"
	"github.com/weareinit/Opal/internal/tools"
)

func GetToken(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return nil, fmt.Errorf("missing token: %w", err)
	}

	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(config.LoadEnv().JWTSecret)
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return token, nil
}

func GetTokenString(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", fmt.Errorf("missing token: %w", err)
	}
	return cookie.Value, nil
}

func GetUserId(r *http.Request) (string, error) {
	token, err := GetToken(r)
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

func GetUser(r *http.Request) (api.User, error) {
    userId, err := GetUserId(r)
    if err != nil {
        return api.User{}, fmt.Errorf("could not get the token from user")
    }

    return tools.LoadDB(func(db *sql.DB) (api.User, error) {
        var user api.User
        getUserQuery := `SELECT user_id, name, family_name, email, discord_username, hacker_id, is_admin, is_volunteer, is_mentor, is_sponsor FROM user WHERE user_id = ?`

        row := db.QueryRow(getUserQuery, userId)

        err := row.Scan(
            &user.UserId,
            &user.Name,
            &user.FamilyName,
            &user.Email,
            &user.DiscordUsername,
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

func IsUserAdmin(userId string, r *http.Request) bool {
  getUser, err := GetUser(r)

  if err != nil {
    return false
  }

  if getUser.IsAdmin {
    return true
  } else {
    return false
  }
}

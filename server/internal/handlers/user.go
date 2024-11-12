package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/weareinit/Opal/internal/helpers"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.GetTokenString(r)
	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	userId, err := helpers.GetUserId(r)
	if err != nil {
		http.Error(w, "Could not retrieve userId", http.StatusUnauthorized)
		return
	}

	// todo: remove this and only send the user's info for the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":  token,
		"userId": userId,
	})
}

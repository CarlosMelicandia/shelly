package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/weareinit/Opal/internal/helpers/user"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := user.GetUser(w, r)

	if err != nil {
		http.Error(w, "Could not fetch the users information", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

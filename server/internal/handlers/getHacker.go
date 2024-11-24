package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/weareinit/Opal/internal/helpers/hacker"
)

func GetHackerHandler(w http.ResponseWriter, r *http.Request) {
	hacker, err := hacker.GetHacker(w, r)

	if err != nil {
		http.Error(w, "Could not fetch the hackers information", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hacker)
}

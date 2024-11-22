package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/cmd/operations"
	"github.com/weareinit/Opal/internal/helpers/user"
	"github.com/weareinit/Opal/internal/utils"
)

func CreateHackerHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := user.GetUserId(w, r)
	if err != nil {
		http.Error(w, `{"error": "Failed to get user ID: `+err.Error()+`"}`, http.StatusUnauthorized)
		return
	}

	// Parse and validate the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error": "Failed to read request body: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	// Parse the request body into the Hacker struct
	hacker, err := parseAndValidateHacker(body)
	if err != nil {
		http.Error(w, `{"error": "Failed to parse and validate request body: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	hacker.UserId = userId

	if _, err := operations.CreateHacker(w, r, hacker); err != nil {
		http.Error(w, `{"error": "Could not create the hacker record: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// Send a JSON response with the redirect URL
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message":  "Hacker created successfully",
		"redirect": "http://localhost:8000/dashboard",
	}
	json.NewEncoder(w).Encode(response)
}

func parseAndValidateHacker(body []byte) (api.Hacker, error) {
	var rawBody map[string]interface{}
	if err := json.Unmarshal(body, &rawBody); err != nil {
		return api.Hacker{}, fmt.Errorf("failed to parse request body as JSON: %v", err)
	}

	if err := utils.ConvertToIntField(&rawBody, "age"); err != nil {
		return api.Hacker{}, fmt.Errorf("invalid age format. Must be an integer: %v", err)
	}
	if err := utils.ConvertToIntField(&rawBody, "grad_year"); err != nil {
		return api.Hacker{}, fmt.Errorf("invalid grad_year format. Must be an integer: %v", err)
	}

	// Convert the updated map back to the Hacker struct
	var hacker api.Hacker
	convertedBody, err := json.Marshal(rawBody)
	if err != nil {
		return api.Hacker{}, fmt.Errorf("failed to re-marshal JSON: %v", err)
	}
	if err := json.Unmarshal(convertedBody, &hacker); err != nil {
		return api.Hacker{}, fmt.Errorf("failed to parse converted request body into struct: %v", err)
	}

	return hacker, nil
}

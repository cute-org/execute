package auth

import (
	"encoding/json"
	"net/http"
)

type ValidateResponse struct {
	Message string `json:"message"`
	User    string `json:"user"`
}

// ValidateHandler handles the /validate GET endpoint to check if the session is valid
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the session token from cookies
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No session token found", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error retrieving session token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the token from the cookie
	token := cookie.Value

	// Check if the session is valid
	username, exists := GetSessionUsername(token)
	if !exists {
		http.Error(w, "Invalid or expired session token", http.StatusUnauthorized)
		return
	}

	// If session is valid, respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ValidateResponse{
		Message: "Session is valid",
		User:    username,
	})
}

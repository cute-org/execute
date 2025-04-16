package user

import (
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"execute/internal"
)

// ServeAvatarHandler handles the /avatar GET endpoint
func ServeAvatarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	var avatarData []byte
	err = internal.DB.QueryRow("SELECT avatar FROM users WHERE id = $1", id).Scan(&avatarData)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not retrieve avatar: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if len(avatarData) == 0 {
		http.Error(w, "No avatar set", http.StatusNotFound)
		return
	}

	encoded := base64.StdEncoding.EncodeToString(avatarData)
	dataURL := "data:image/png;base64," + encoded

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(dataURL)); err != nil {
		log.Printf("Failed to write base64 avatar: %v", err)
	}
}

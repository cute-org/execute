package user

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"execute/internal"
)

type AvatarResponse struct {
	Avatar string `json:"avatar"`
}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AvatarResponse{Avatar: dataURL})
}

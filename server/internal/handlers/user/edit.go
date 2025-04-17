package user

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"execute/internal"
	"execute/internal/handlers/auth"
)

type EditUserRequest struct {
	ID          int     `json:"id"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`    // Current password
	NewPassword *string `json:"newpassword,omitempty"` // New password
	Avatar      *string `json:"avatar,omitempty"`
}

type EditUserResponse struct {
	Status string `json:"status"`
}

// EditUserHandler handles the /user PUT endpoint
func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		handleJSONUpdate(w, r)
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		handleMultipartUpdate(w, r)
	} else {
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
	}
}

// JSON handling
func handleJSONUpdate(w http.ResponseWriter, r *http.Request) {
	var req EditUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.ID == 0 || req.Password == nil {
		http.Error(w, "User ID and current password are required", http.StatusBadRequest)
		return
	}

	// Retrieve the user's stored password hash and salt from the database
	var storedPasswordHash string
	var storedSalt string
	err := internal.DB.QueryRow("SELECT passwordhash, salt FROM users WHERE id = $1", req.ID).Scan(&storedPasswordHash, &storedSalt)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert storedSalt (string) to []byte before passing it to HashPassword
	saltBytes, err := auth.DecodeSalt(storedSalt)
	if err != nil {
		http.Error(w, "Server error decoding salt", http.StatusInternalServerError)
		return
	}

	// Verify the provided password against the stored hash and salt
	passwordHash := auth.HashPassword(*req.Password, saltBytes)
	if passwordHash != storedPasswordHash {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	// Build dynamic update
	var updates []string
	var args []interface{}
	argPos := 1

	if req.Username != nil {
		updates = append(updates, "username = $"+strconv.Itoa(argPos))
		args = append(args, *req.Username)
		argPos++
	}

	// Handle password change if newpassword is provided
	if req.NewPassword != nil {
		if len(*req.NewPassword) < 8 {
			http.Error(w, "New password is too short", http.StatusBadRequest)
			return
		}

		salt, err := auth.GenerateSalt()
		if err != nil {
			http.Error(w, "Failed to generate salt", http.StatusInternalServerError)
			return
		}

		newPasswordHash := auth.HashPassword(*req.NewPassword, salt)
		updates = append(updates, "salt = $"+strconv.Itoa(argPos))
		args = append(args, auth.EncodeSalt(salt))
		argPos++
		updates = append(updates, "passwordhash = $"+strconv.Itoa(argPos))
		args = append(args, newPasswordHash)
		argPos++
	}

	// Handle avatar update if provided
	if req.Avatar != nil {
		updates = append(updates, "avatar = $"+strconv.Itoa(argPos))
		args = append(args, *req.Avatar)
		argPos++
	}

	if len(updates) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	args = append(args, req.ID)
	query := "UPDATE users SET " + strings.Join(updates, ", ") + " WHERE id = $" + strconv.Itoa(argPos)

	_, err = internal.DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := EditUserResponse{Status: "updated"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Multipart (file) upload handling
func handleMultipartUpdate(w http.ResponseWriter, r *http.Request) {
	// Limit request body to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10MB

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Could not parse form or file too big: "+err.Error(), http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(idStr)

	username := r.FormValue("username")
	password := r.FormValue("password")

	var avatarBytes []byte
	file, _, err := r.FormFile("avatar")
	if err == nil {
		defer file.Close()
		avatarBytes, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read avatar: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if len(avatarBytes) > 10<<20 {
			http.Error(w, "Avatar file exceeds 10MB", http.StatusRequestEntityTooLarge)
			return
		}
	}

	// Build dynamic update
	var updates []string
	var args []interface{}
	argPos := 1

	if username != "" {
		updates = append(updates, "username = $"+strconv.Itoa(argPos))
		args = append(args, username)
		argPos++
	}
	if password != "" {
		salt, err := auth.GenerateSalt()
		if err != nil {
			http.Error(w, "Failed to generate salt", http.StatusInternalServerError)
			return
		}
		hash := auth.HashPassword(password, salt)
		updates = append(updates, "salt = $"+strconv.Itoa(argPos))
		args = append(args, auth.EncodeSalt(salt))
		argPos++
		updates = append(updates, "passwordhash = $"+strconv.Itoa(argPos))
		args = append(args, hash)
		argPos++
	}
	if avatarBytes != nil {
		updates = append(updates, "avatar = $"+strconv.Itoa(argPos))
		args = append(args, avatarBytes)
		argPos++
	}

	if len(updates) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	args = append(args, id)
	query := "UPDATE users SET " + strings.Join(updates, ", ") + " WHERE id = $" + strconv.Itoa(argPos)

	_, err = internal.DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := EditUserResponse{Status: "updated"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

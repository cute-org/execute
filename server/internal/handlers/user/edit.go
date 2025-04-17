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
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		handleJSONUpdate(w, r, userID)
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		handleMultipartUpdate(w, r, userID)
	} else {
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
	}
}

// JSON handling
func handleJSONUpdate(w http.ResponseWriter, r *http.Request, userID int) {
	var req EditUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.Password == nil {
		http.Error(w, "Current password is required", http.StatusBadRequest)
		return
	}

	// Retrieve the user's stored password hash and salt from the database
	var storedHash, storedSaltStr string
	if err := internal.DB.QueryRow(
		"SELECT passwordhash, salt FROM users WHERE id=$1", userID,
	).Scan(&storedHash, &storedSaltStr); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert storedSalt (string) to []byte before passing it to HashPassword
	saltBytes, err := auth.DecodeSalt(storedSaltStr)
	if err != nil {
		http.Error(w, "Server error decoding salt", http.StatusInternalServerError)
		return
	}
	if auth.HashPassword(*req.Password, saltBytes) != storedHash {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	var updates []string
	var args []any
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
		newSalt, err := auth.GenerateSalt()
		if err != nil {
			http.Error(w, "Failed to generate salt", http.StatusInternalServerError)
			return
		}
		updates = append(updates, "salt = $"+strconv.Itoa(argPos))
		args = append(args, auth.EncodeSalt(newSalt))
		argPos++
		updates = append(updates, "passwordhash = $"+strconv.Itoa(argPos))
		args = append(args, auth.HashPassword(*req.NewPassword, newSalt))
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

	args = append(args, userID)
	query := "UPDATE users SET " + strings.Join(updates, ", ") + " WHERE id = $" + strconv.Itoa(argPos)
	if _, err := internal.DB.Exec(query, args...); err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EditUserResponse{Status: "updated"})
}

// Multipart (file) upload handling
func handleMultipartUpdate(w http.ResponseWriter, r *http.Request, userID int) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Could not parse form or file too big: "+err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var avatarBytes []byte
	if file, _, err := r.FormFile("avatar"); err == nil {
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

	var updates []string
	var args []any
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

	args = append(args, userID)
	query := "UPDATE users SET " + strings.Join(updates, ", ") + " WHERE id = $" + strconv.Itoa(argPos)
	if _, err := internal.DB.Exec(query, args...); err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EditUserResponse{Status: "updated"})
}

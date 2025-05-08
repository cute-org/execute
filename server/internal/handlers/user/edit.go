package user

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
)

type EditUserRequest struct {
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`    // Current password
	NewPassword *string `json:"newpassword,omitempty"` // New password
	Avatar      *string `json:"avatar,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	BirthDate   *string `json:"birth_date,omitempty"`
	Role        *string `json:"role,omitempty"`
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
		base64Data := *req.Avatar

		switch {
		case strings.HasPrefix(base64Data, "data:image/png;base64,"):
			base64Data = strings.TrimPrefix(base64Data, "data:image/png;base64,")
		case strings.HasPrefix(base64Data, "data:image/jpeg;base64,"):
			base64Data = strings.TrimPrefix(base64Data, "data:image/jpeg;base64,")
		case strings.HasPrefix(base64Data, "data:image/jpg;base64,"):
			base64Data = strings.TrimPrefix(base64Data, "data:image/jpg;base64,")
		default:
			http.Error(w, "Unsupported image format; only PNG, JPEG and JPG are allowed", http.StatusBadRequest)
			return
		}

		avatarBytes, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			http.Error(w, "Invalid base64 avatar data", http.StatusBadRequest)
			return
		}
		updates = append(updates, "avatar = $"+strconv.Itoa(argPos))
		args = append(args, avatarBytes)
		argPos++
	}

	if req.DisplayName != nil {
		updates = append(updates, "display_name = $"+strconv.Itoa(argPos))
		args = append(args, *req.DisplayName)
		argPos++
	}
	if req.Phone != nil {
		updates = append(updates, "phone = $"+strconv.Itoa(argPos))
		args = append(args, *req.Phone)
		argPos++
	}
	if req.BirthDate != nil {
		// Expect YYYY-MM-DD format
		if _, err := time.Parse("2006-01-02", *req.BirthDate); err != nil {
			http.Error(w, "Invalid birth_date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		updates = append(updates, "birth_date = $"+strconv.Itoa(argPos))
		args = append(args, *req.BirthDate)
		argPos++
	}
	if req.Role != nil {
		updates = append(updates, "role = $"+strconv.Itoa(argPos))
		args = append(args, *req.Role)
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

	// Form values
	username := r.FormValue("username")
	password := r.FormValue("password")
	displayName := r.FormValue("display_name")
	phone := r.FormValue("phone")
	birthDate := r.FormValue("birth_date")
	role := r.FormValue("role")

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
	if displayName != "" {
		updates = append(updates, "display_name = $"+strconv.Itoa(argPos))
		args = append(args, displayName)
		argPos++
	}
	if phone != "" {
		updates = append(updates, "phone = $"+strconv.Itoa(argPos))
		args = append(args, phone)
		argPos++
	}
	if birthDate != "" {
		// Expect YYYY-MM-DD format
		if _, err := time.Parse("2006-01-02", birthDate); err != nil {
			http.Error(w, "Invalid birth_date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		updates = append(updates, "birth_date = $"+strconv.Itoa(argPos))
		args = append(args, birthDate)
		argPos++
	}
	if role != "" {
		updates = append(updates, "role = $"+strconv.Itoa(argPos))
		args = append(args, role)
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

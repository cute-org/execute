package user

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"

	"execute/internal"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// UsersHandler handles the /users GET endpoint
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := internal.DB.Query("SELECT id, username FROM users")
	if err != nil {
		http.Error(w, "Failed to query users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			http.Error(w, "Failed to scan user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Check for any errors that occurred during the iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header for JSON response
	w.WriteHeader(http.StatusOK)
	if len(users) == 0 {
		w.Write([]byte("[]"))
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

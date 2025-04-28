package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"execute/internal"
	"execute/internal/handlers/auth"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type GroupMember struct {
	ID          int        `json:"id"`
	Username    string     `json:"username"`
	Role        string     `json:"role"`
	DisplayName *string    `json:"display_name,omitempty"`
	Phone       *string    `json:"phone,omitempty"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
}

// UsersHandler handles the /user GET endpoint
func UsersHandler(w http.ResponseWriter, r *http.Request) {
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

// GroupUsersHanlder handles the /group GET endpoint
func GroupUsersHanlder(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	groupID, err := GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "You are not in a group: "+err.Error(), http.StatusForbidden)
		return
	}

	rows, err := internal.DB.Query(`
    SELECT id, username, role, display_name, phone, birth_date
      FROM users
     WHERE group_id = $1
  `, groupID)
	if err != nil {
		http.Error(w, "Failed to query group members: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	members := make([]GroupMember, 0)
	for rows.Next() {
		var m GroupMember
		var dn, ph sql.NullString
		var bd sql.NullTime

		if err := rows.Scan(
			&m.ID,
			&m.Username,
			&m.Role,
			&dn,
			&ph,
			&bd,
		); err != nil {
			http.Error(w, "Failed to scan member: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// NULL-handling
		if dn.Valid {
			m.DisplayName = &dn.String
		}
		if ph.Valid {
			m.Phone = &ph.String
		}
		if bd.Valid {
			m.BirthDate = &bd.Time
		}

		members = append(members, m)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error during rows iteration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if len(members) == 0 {
		w.Write([]byte("[]"))
	} else {
		json.NewEncoder(w).Encode(members)
	}
}

package group

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"execute/internal"
	"execute/internal/handlers/auth"
)

type createReq struct {
	Name string `json:"name"`
}

type createResp struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type joinReq struct {
	Code string `json:"code"`
}

type resp struct {
	Message string `json:"message"`
}

type updateGroupReq struct {
	GroupID int    `json:"groupId"`
	Name    string `json:"name"`
}

// CreateGroupHandler handles POST /group
func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "Group name is required", http.StatusBadRequest)
		return
	}

	code, err := NewCode()
	if err != nil {
		http.Error(w, "Failed to generate code", http.StatusInternalServerError)
		return
	}

	var id int
	err = internal.DB.QueryRow(
		`INSERT INTO groups(name, code, creator_user_id)
         VALUES($1,$2,$3) RETURNING id`,
		req.Name, code, userID,
	).Scan(&id)
	if err != nil {
		http.Error(w, "Could not create group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createResp{ID: id, Code: code})
}

// JoinGroupHandler handles POST /group/join
func JoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req joinReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Code == "" {
		http.Error(w, "Join code is required", http.StatusBadRequest)
		return
	}

	var existing sql.NullInt64
	err = internal.DB.QueryRow(
		"SELECT group_id FROM users WHERE id = $1", userID,
	).Scan(&existing)
	if err != nil {
		http.Error(w, "User lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if existing.Valid {
		http.Error(w, "You are already in a group", http.StatusConflict)
		return
	}

	var groupID int
	err = internal.DB.QueryRow(
		"SELECT id FROM groups WHERE code = $1", req.Code,
	).Scan(&groupID)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid join code", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Lookup error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = internal.DB.Exec(
		"UPDATE users SET group_id = $1 WHERE id = $2",
		groupID, userID,
	)
	if err != nil {
		http.Error(w, "Could not join group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp{Message: "Joined group successfully"})
}

// UpdateGroupHandler handles PUT /group
func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req updateGroupReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "New group name is required", http.StatusBadRequest)
		return
	}

	var creatorID int
	err = internal.DB.QueryRow(
		"SELECT creator_user_id FROM groups WHERE id=$1", req.GroupID,
	).Scan(&creatorID)
	if err == sql.ErrNoRows {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if creatorID != userID {
		http.Error(w, "Forbidden: only the creator can edit", http.StatusForbidden)
		return
	}

	_, err = internal.DB.Exec(
		"UPDATE groups SET name=$1 WHERE id=$2",
		req.Name, req.GroupID,
	)
	if err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message":"Group updated successfully"}`)
}

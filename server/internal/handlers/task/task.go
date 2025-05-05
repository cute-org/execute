package task

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
	"execute/internal/handlers/user"
)

type Task struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"groupId"`
	CreatorUserID   int       `json:"creatorUserId"`
	CreatorUsername string    `json:"creatorUsername"`
	CreationDate    time.Time `json:"creationDate"`
	DueDate         time.Time `json:"dueDate"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	PointsValue     int       `json:"pointsValue"`
	Step            int       `json:"step"`
}

type createReq struct {
	DueDate     time.Time `json:"dueDate"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsValue int       `json:"pointsValue"`
	Step        int       `json:"step"`
}

type updateTaskReq struct {
	TaskID      int       `json:"taskId"`
	DueDate     time.Time `json:"dueDate"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsValue int       `json:"pointsValue"`
}

type StepUpdateReq struct {
	TaskID int    `json:"taskId"`
	Action string `json:"action"`
}

// CreateTaskHandler handles POST /task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.PointsValue < 0 {
		http.Error(w, "Name required and points must be ≥0", http.StatusBadRequest)
		return
	}

	var taskID int
	err = internal.DB.QueryRow(
		`INSERT INTO tasks
		   (group_id, creator_user_id, due_date, name, description, points_value, step)
		 VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		groupID, userID, req.DueDate, req.Name, req.Description, req.PointsValue, req.Step,
	).Scan(&taskID)
	if err != nil {
		http.Error(w, "Failed to create task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var username string
	if err := internal.DB.QueryRow(
		"SELECT username FROM users WHERE id=$1", userID,
	).Scan(&username); err != nil {
		http.Error(w, "Failed to retrieve username: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id":              taskID,
		"creatorUsername": username,
	})
}

// ListTasksHandler handles GET /task
func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	rows, err := internal.DB.Query(
		`SELECT
		  t.id,
		  t.group_id,
		  t.creator_user_id,
		  u.username,
		  t.creation_date,
		  t.due_date,
		  t.name,
		  t.description,
		  t.points_value,
		  t.step
		FROM tasks t
		JOIN users u ON u.id = t.creator_user_id
		WHERE t.group_id = $1`,
		groupID,
	)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(
			&t.ID,
			&t.GroupID,
			&t.CreatorUserID,
			&t.CreatorUsername,
			&t.CreationDate,
			&t.DueDate,
			&t.Name,
			&t.Description,
			&t.PointsValue,
			&t.Step,
		); err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTaskHandler handles PUT /task
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req updateTaskReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.PointsValue < 0 {
		http.Error(w, "Name required and points must be ≥0", http.StatusBadRequest)
		return
	}

	var creatorID int
	err = internal.DB.QueryRow(
		"SELECT creator_user_id FROM tasks WHERE id=$1", req.TaskID,
	).Scan(&creatorID)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
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
		`UPDATE tasks
		    SET name=$1,
		        description=$2,
		        due_date=$3,
		        points_value=$4
		  WHERE id=$5`,
		req.Name, req.Description, req.DueDate, req.PointsValue, req.TaskID,
	)
	if err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message":"Task updated successfully"}`)
}

// UpdateStepHandler handles PATCH /task
func TaskStepHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var req StepUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the action
	if req.Action != "+1" && req.Action != "-1" {
		http.Error(w, "Invalid action. Must be '+1' or '-1'", http.StatusBadRequest)
		return
	}

	// Get the userID from the auth header
	userID, err := auth.GetUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Retrieve the groupID for the user
	groupID, err := user.GetUserGroupID(userID)
	if err != nil {
		http.Error(w, "Group lookup failed: "+err.Error(), http.StatusForbidden)
		return
	}

	// Retrieve the task's group ID to ensure the user belongs to the same group
	var taskGroupID, currentStep int
	err = internal.DB.QueryRow(
		"SELECT group_id, step FROM tasks WHERE id=$1", req.TaskID,
	).Scan(&taskGroupID, &currentStep)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Task lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Ensure the user is part of the same group
	if taskGroupID != groupID {
		http.Error(w, "Forbidden: You are not in the same group as the task", http.StatusForbidden)
		return
	}

	// Determine the step update
	var stepChange int
	if req.Action == "+1" {
		stepChange = 1
	} else if req.Action == "-1" {
		stepChange = -1
	}

	// Update the step of the task in the database
	_, err = internal.DB.Exec(
		`UPDATE tasks
		 SET step = step + $1
		 WHERE id = $2`,
		stepChange, req.TaskID,
	)
	if err != nil {
		http.Error(w, "Failed to update step: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the updated step value
	err = internal.DB.QueryRow(
		"SELECT step FROM tasks WHERE id=$1", req.TaskID,
	).Scan(&currentStep)
	if err != nil {
		http.Error(w, "Failed to retrieve updated step value: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated task information
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"taskId":  req.TaskID,
		"step":    currentStep,
		"message": "Task step updated successfully",
	})
}

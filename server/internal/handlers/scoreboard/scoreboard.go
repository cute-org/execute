package scoreboard

import (
	"encoding/json"
	"net/http"

	"execute/internal"
)

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PointsScore int    `json:"points_score"`
}

// ScoreboardHandler handles GET /scoreboard
func ScoreboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Query groups sorted by points_score
	rows, err := internal.DB.Query(`
		SELECT id, name, points_score
		FROM groups
		ORDER BY points_score DESC, id ASC
	`)
	if err != nil {
		http.Error(w, "failed to query groups: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var g Group
		if err := rows.Scan(&g.ID, &g.Name, &g.PointsScore); err != nil {
			http.Error(w, "failed to scan group: "+err.Error(), http.StatusInternalServerError)
			return
		}
		groups = append(groups, g)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "rows iteration error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

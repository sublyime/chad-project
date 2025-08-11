package api

import (
	"encoding/json"
	"net/http"
)

// Health check endpoint
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "CHAD API running"})
}

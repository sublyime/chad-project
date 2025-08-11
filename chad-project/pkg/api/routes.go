package api

import (
	"encoding/json"
	"net/http"
)

// Health check
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "CHAD API running"})
}

// Placeholder chemical search
func ChemicalHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate returning a chemical
	json.NewEncoder(w).Encode(map[string]interface{}{
		"name":   "Chlorine",
		"cas":    "7782-50-5",
		"mw":     70.9,
		"bp":     -34.04,
		"hazard": "Toxic gas",
	})
}

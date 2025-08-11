package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sublyime/chad-project/pkg/api"
	"github.com/sublyime/chad-project/pkg/db"
)

func main() {
	// Connect DB
	db.ConnectDB()

	r := chi.NewRouter()

	// API routes
	r.Get("/api/health", api.HealthHandler)
	r.Get("/api/chemicals", api.ChemicalHandler)

	log.Println("✅ CHAD backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

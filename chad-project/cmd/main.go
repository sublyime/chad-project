package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sublyime/chad-project/chad-project/pkg/db"
	// replace 'your_module_path' with your actual module path
)

func main() {
	// Initialize the database connection
	db.ConnectDB()

	// Create a new router
	r := chi.NewRouter()

	// Define a simple root endpoint for testing
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to CHAD! Database connected successfully."))
	})

	// TODO: Add routes for API, weather, chemical queries, reports, etc.

	// Start the HTTP server on port 8080
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

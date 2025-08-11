package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectDB() {
	server := getEnv("DB_SERVER", "localhost")
	port := getEnv("DB_PORT", "5432") // default Postgres port
	user := getEnv("DB_USER", "chad1")
	password := getEnv("DB_PASSWORD", "chad")
	database := getEnv("DB_NAME", "chad-project")

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server, port, user, password, database,
	)

	var err error
	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error creating connection pool: %s", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	log.Println("✅ Connected to PostgreSQL database:", database)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

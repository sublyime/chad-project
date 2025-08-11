package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // Register SQL Server driver
)

var DB *sql.DB

// ConnectDB initializes the connection to the SQL Express database
func ConnectDB() {
	// Defaults set as requested — can be overridden via env vars
	server := getEnv("DB_SERVER", `SUBMAIN\SQLEXPRESS`)
	port := getEnv("DB_PORT", "1433")
	user := getEnv("DB_USER", "chad")
	password := getEnv("DB_PASSWORD", "chad")
	database := getEnv("DB_NAME", "chad-project")

	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable",
		server, user, password, port, database,
	)

	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error creating connection pool: %s", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	log.Println("✅ Connected to SQL Express database:", database)
}

// getEnv returns an environment variable or a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

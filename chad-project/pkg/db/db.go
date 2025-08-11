package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // Register driver
)

var DB *sql.DB

func ConnectDB() {
	// Connect directly to server without instance name; use fixed TCP port 1433
	server := getEnv("DB_SERVER", "SUBMAIN") // remove \SQLEXPRESS instance name
	port := getEnv("DB_PORT", "1433")        // fixed port configured in SQL Server
	user := getEnv("DB_USER", "chad1")
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

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

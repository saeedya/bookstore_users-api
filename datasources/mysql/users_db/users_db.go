package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for Go
)

var (
	Client *sql.DB
)

func init() {
	// Fetch database credentials from environment variables
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Database environment variables not set")
	}

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open a connection to the database
	Client, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Close()

	// Verify the connection is successful
	if err := Client.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

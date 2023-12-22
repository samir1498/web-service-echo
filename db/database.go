// db/db.go

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve values from environment variables or use default values
	host := getEnv("DB_HOST")
	port := getEnv("DB_PORT")
	user := getEnv("DB_USER")
	password := getEnv("DB_PASSWORD")
	dbname := getEnv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var errDB error
	DB, errDB = sql.Open("postgres", psqlInfo)
	if errDB != nil {
		panic(errDB)
	}

	errDB = DB.Ping()
	if errDB != nil {
		panic(errDB)
	}

	fmt.Println("Successfully connected to the database!")

	createTableIfNotExists()
}

func createTableIfNotExists() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

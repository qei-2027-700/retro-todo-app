package storage

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
	log.Println("[DB] Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	log.Printf("[DB] Connecting to database: host=%s port=%s user=%s dbname=%s", dbHost, dbPort, dbUser, dbName)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		os.Getenv("DB_PASSWORD"),
		dbName,
	)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[DB] Failed to open database connection: %v", err)
	}

	log.Println("[DB] Pinging database...")
	if err = DB.Ping(); err != nil {
		log.Fatalf("[DB] Failed to ping database: %v", err)
	}
	log.Println("[DB] Successfully connected to database!")
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("[MIGRATION] Starting database migration...")

	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// データベース接続
	db := connectDB()
	defer db.Close()

	// マイグレーション履歴テーブルの確認・作成
	ensureSchemaVersionTable(db)

	// 適用済みのマイグレーションを取得
	appliedMigrations := getAppliedMigrations(db)
	appliedMap := make(map[int]bool)
	for _, version := range appliedMigrations {
		appliedMap[version] = true
	}

	// マイグレーションファイルの取得とソート
	migrationFiles, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		log.Fatalf("[MIGRATION] Could not find migration files: %v", err)
	}
	sort.Strings(migrationFiles)

	log.Println("[MIGRATION] Found migration files:", migrationFiles)

	// マイグレーションの実行
	for _, file := range migrationFiles {
		version, err := getVersionFromFile(file)
		if err != nil {
			log.Printf("[MIGRATION] Skipping file with invalid version format: %s", file)
			continue
		}

		if _, applied := appliedMap[version]; applied {
			log.Printf("[MIGRATION] Skipping already applied migration: %s", file)
			continue
		}

		log.Printf("[MIGRATION] Applying migration: %s", file)
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("[MIGRATION] Failed to read migration file %s: %v", file, err)
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("[MIGRATION] Failed to begin transaction: %v", err)
		}

		if _, err := tx.Exec(string(content)); err != nil {
			tx.Rollback()
			log.Fatalf("[MIGRATION] Failed to apply migration %s: %v", file, err)
		}

		if err := addMigrationHistory(tx, version); err != nil {
			tx.Rollback()
			log.Fatalf("[MIGRATION] Failed to record migration history for version %d: %v", version, err)
		}

		if err := tx.Commit(); err != nil {
			log.Fatalf("[MIGRATION] Failed to commit transaction for migration %s: %v", file, err)
		}

		log.Printf("[MIGRATION] ✓ Successfully applied migration: %s", file)
	}

	log.Println("[MIGRATION] Migration completed successfully!")
}

func connectDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[MIGRATION] Failed to connect to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("[MIGRATION] Failed to ping database: %v", err)
	}
	log.Println("[MIGRATION] Successfully connected to database!")
	return db
}

func ensureSchemaVersionTable(db *sql.DB) {
	log.Println("[MIGRATION] Ensuring schema_migrations table exists...")
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INT PRIMARY KEY
		);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("[MIGRATION] Failed to create or verify schema_migrations table: %v", err)
	}
	log.Println("[MIGRATION] ✓ schema_migrations table verified.")
}

func getAppliedMigrations(db *sql.DB) []int {
	rows, err := db.Query("SELECT version FROM schema_migrations ORDER BY version")
	if err != nil {
		log.Fatalf("[MIGRATION] Failed to query schema_migrations table: %v", err)
	}
	defer rows.Close()

	var versions []int
	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			log.Fatalf("[MIGRATION] Failed to scan migration version: %v", err)
		}
		versions = append(versions, version)
	}
	return versions
}

func getVersionFromFile(file string) (int, error) {
	base := filepath.Base(file)
	parts := strings.Split(base, "_")
	if len(parts) < 1 {
		return 0, fmt.Errorf("invalid migration file name format: %s", base)
	}
	return strconv.Atoi(parts[0])
}

func addMigrationHistory(tx *sql.Tx, version int) error {
	query := "INSERT INTO schema_migrations (version) VALUES ($1)"
	_, err := tx.Exec(query, version)
	return err
}

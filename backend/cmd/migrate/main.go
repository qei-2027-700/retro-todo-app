package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("[MIGRATION] Starting database migration...")

	// .envファイルの読み込み
	log.Println("[MIGRATION] Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// データベース接続情報の取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	log.Printf("[MIGRATION] Connecting to database: host=%s port=%s user=%s dbname=%s", dbHost, dbPort, dbUser, dbName)

	// データベース接続
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("[MIGRATION] Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 接続確認
	if err = db.Ping(); err != nil {
		log.Fatalf("[MIGRATION] Failed to ping database: %v", err)
	}
	log.Println("[MIGRATION] Successfully connected to database!")

	// マイグレーション実行
	log.Println("[MIGRATION] Running migrations...")

	// todosテーブルの作成
	createTodosTable := `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			completed BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	log.Println("[MIGRATION] Creating todos table if not exists...")
	_, err = db.Exec(createTodosTable)
	if err != nil {
		log.Fatalf("[MIGRATION] Failed to create todos table: %v", err)
	}
	log.Println("[MIGRATION] ✓ todos table created/verified")

	// 既存テーブルへのカラム追加（IF NOT EXISTSで冪等性を保つ）
	log.Println("[MIGRATION] Adding columns if not exists...")

	alterTableQueries := []string{
		"ALTER TABLE todos ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
		"ALTER TABLE todos ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
	}

	for _, query := range alterTableQueries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatalf("[MIGRATION] Failed to alter table: %v", err)
		}
	}
	log.Println("[MIGRATION] ✓ Columns added/verified")

	// インデックスの作成
	createIndexes := `
		CREATE INDEX IF NOT EXISTS idx_todos_completed ON todos(completed);
	`

	log.Println("[MIGRATION] Creating indexes if not exists...")
	_, err = db.Exec(createIndexes)
	if err != nil {
		log.Fatalf("[MIGRATION] Failed to create indexes: %v", err)
	}
	log.Println("[MIGRATION] ✓ Indexes created/verified")

	log.Println("[MIGRATION] Migration completed successfully!")
}

package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/storage"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("[MAIN] Starting server initialization...")
	storage.Init()

	// リポジトリの初期化
	todoRepo := repository.NewTodoRepository(storage.DB)

	// ハンドラーの初期化
	todoHandler := handler.NewTodoHandler(todoRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("[MAIN] Registering routes...")
	e.GET("/todos", todoHandler.GetTodos)
	e.POST("/todos", todoHandler.CreateTodo)

	log.Println("[MAIN] Server starting on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/storage"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "backend/docs" // swagger docs
)

// @title Retro Todo API
// @version 1.0
// @description レトロなTODOアプリケーションのAPI
// @host localhost:8080
// @BasePath /
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
	e.PUT("/todos", todoHandler.UpdateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Println("[MAIN] Server starting on :8080")
	log.Println("[MAIN] Swagger UI: http://localhost:8080/swagger/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}

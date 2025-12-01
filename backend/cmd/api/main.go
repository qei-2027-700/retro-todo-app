package main

import (
	"backend/internal/handler"
	authmw "backend/internal/middleware"
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
	sprintRepo := repository.NewSprintRepository(storage.DB)
	userRepo := repository.NewUserRepository(storage.DB)

	// ハンドラーの初期化
	todoHandler := handler.NewTodoHandler(todoRepo)
	sprintHandler := handler.NewSprintHandler(sprintRepo)
	authHandler := handler.NewAuthHandler(userRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 認証不要エンドポイント
	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)

	// Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 認証必要エンドポイント
	protected := e.Group("")
	protected.Use(authmw.AuthMiddleware())

	// todos
	protected.GET("/todos", todoHandler.GetTodos)
	protected.POST("/todos", todoHandler.CreateTodo)
	protected.POST("/todos/search", todoHandler.SearchTodos)
	protected.PUT("/todos/:id", todoHandler.UpdateTodo)
	protected.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// sprints
	protected.GET("/sprints", sprintHandler.GetSprints)
	protected.POST("/sprints", sprintHandler.CreateSprint)
	protected.POST("/sprints/search", sprintHandler.SearchSprints)
	protected.PUT("/sprints/:id", sprintHandler.UpdateSprint)
	protected.DELETE("/sprints/:id", sprintHandler.DeleteSprint)

	log.Println("[MAIN] Server starting on :8080")
	log.Println("[MAIN] Swagger UI: http://localhost:8080/swagger/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}

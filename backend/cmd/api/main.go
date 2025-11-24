package main

import (
	"backend/internal/handler"
	"backend/internal/storage"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("[MAIN] Starting server initialization...")
	storage.Init()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("[MAIN] Registering routes...")
	e.GET("/todos", handler.GetTodos)
	e.POST("/todos", handler.CreateTodo)

	log.Println("[MAIN] Server starting on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

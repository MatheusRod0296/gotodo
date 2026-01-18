package main

import (
	"go-todo/internal/config"
	"go-todo/internal/database"
	"go-todo/internal/sorterUrl"
	"go-todo/internal/todo"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Open(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	repo := todo.NewRepository()
	service := todo.NewService(repo)
	todoHandler := todo.NewHandler(service)

	urlRepo := sorterUrl.NewURLRepository(db)
	urlService := sorterUrl.NewURLService(urlRepo)
	handler := sorterUrl.NewHandler(urlService)

	r.GET("/todos", todoHandler.List)
	r.GET("/todos/:id", todoHandler.GetById)
	r.POST("/todos", todoHandler.Create)
	r.PUT("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)

	r.POST("/shorten", handler.Create)
	r.GET("/:code", handler.Redirect)
	r.GET("/list/:offset/:limit", handler.List)

	r.Run(":8080")
}

package main

import (
	"go-todo/internal/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := todo.NewRepository()
	service := todo.NewService(repo)
	todoHandler := todo.NewHandler(service)

	r.GET("/todos", todoHandler.List)
	r.POST("/todos", todoHandler.Create)
	r.PUT("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)

	r.Run(":8080")
}

package main

import (
	"todo/endpoints/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	todoGroup := r.Group("/todo")
	todo.RegisterTodoEndpoints(todoGroup)

	// Run application
	r.Run(":8080")
}

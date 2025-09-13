package main

import (
    "github.com/gin-gonic/gin"
	"todo/endpoints/todo"
)


func main (){
	r := gin.Default()

	// Routes
	todoGroup := r.Group("/todo")
	todo.RegisterTodoEndpoints(todoGroup)
}
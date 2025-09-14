package main

import (
	"github.com/gin-gonic/gin"
	"todo/endpoints/todo"
)

func main() {
	r := gin.Default()

	// create /todo group
	todo.Register(r.Group("/todo"))

	r.Run(":8080")
}

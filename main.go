package main

import (
	"todo/endpoints/todo"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo API
// @version 1.0
// @description This is a sample todo app with Gin
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// register endpoints here...
	todo.Register(r.Group("/todo"))

	// swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

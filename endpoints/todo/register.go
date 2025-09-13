package todo

import "github.com/gin-gonic/gin"

func RegisterTodoEndpoints(rg *gin.RouterGroup) {
	rg.GET("/list", GetTodosList)
	// rg.POST("/upsert", UpsertTodo)
	// rg.DELETE("/delete",DeleteTodo)
}

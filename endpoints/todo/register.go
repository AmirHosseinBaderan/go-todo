package todo;

import "github.com/gin-gonic/gin"

func RegisterTodoEndpoints(rg *gin.RouterGroup){
	rg.POST("/upsert",UpsertTodo)
	rg.GET("/list",GetTodoList)
	rg.DELETE("/delete",DeleteTodo)
}
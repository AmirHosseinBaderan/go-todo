package todo

import (
	"github.com/gin-gonic/gin"
	"todo/endpoints/todo/delete"
	"todo/endpoints/todo/getlist"
	"todo/endpoints/todo/upsert"
)

func Register(rg *gin.RouterGroup) {
	rg.GET("/list", getlist.Handler)
	rg.POST("/upsert", upsert.Handler)
	rg.DELETE("/:id", delete.Handler)
}

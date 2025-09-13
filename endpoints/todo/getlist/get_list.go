package todo

import "github.com/gin-gonic/gin"

func GetTodosList(c *gin.Context) {
	c.JSON(200, gin.H{"todos": todos})
}

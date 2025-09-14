package getlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var todos = []map[string]interface{}{
	{"id": 1, "title": "Learn Gin", "done": false},
	{"id": 2, "title": "Build Todo App", "done": false},
}

func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

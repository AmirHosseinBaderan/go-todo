package delete

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos = []map[string]interface{}{
	{"id": 1, "title": "Learn Gin", "done": false},
	{"id": 2, "title": "Build Todo App", "done": false},
}

func Handler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, todo := range todos {
		if todo["id"] == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

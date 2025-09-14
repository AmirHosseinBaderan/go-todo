package getlist

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/endpoints/todo/store"
)

// @Summary Get todos
// @Description Get the list of all todos
// @Tags todos
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todo/list [get]
func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"todos": store.Todos})
}

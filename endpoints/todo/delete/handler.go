package delete

import (
	"net/http"
	"strconv"
	"todo/endpoints/todo/store"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a todo
// @Description Delete a todo by its ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Produce json
// @Success 200 {object} map[string]string "Todo deleted"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Todo not found"
// @Router /todo/delete/{id} [delete]
func Handler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, todo := range store.Todos {
		if todo.ID == id {
			store.Todos = append(store.Todos[:i], store.Todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

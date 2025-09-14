package upsert

import (
	"net/http"
	"todo/endpoints/todo/model"
	"todo/endpoints/todo/store"

	"github.com/gin-gonic/gin"
)

// @Summary Add or update a todo
// @Description Add a new todo or update an existing one if ID matches
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body model.Todo true "Todo data"
// @Success 201 {object} model.Todo "Todo created"
// @Success 200 {object} model.Todo "Todo updated"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /todo/upsert [post]
func Handler(c *gin.Context) {
	var newTodo model.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if ID exists, update; otherwise create
	for i, t := range store.Todos {
		if t.ID == newTodo.ID {
			store.Todos[i] = newTodo
			c.JSON(http.StatusOK, newTodo)
			return
		}
	}

	newTodo.ID = len(store.Todos) + 1
	store.Todos = append(store.Todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

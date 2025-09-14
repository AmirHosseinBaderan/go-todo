package upsert

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Gin", Done: false},
	{ID: 2, Title: "Build Todo App", Done: false},
}

func Handler(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if ID exists, update; otherwise create
	for i, t := range todos {
		if t.ID == newTodo.ID {
			todos[i] = newTodo
			c.JSON(http.StatusOK, newTodo)
			return
		}
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

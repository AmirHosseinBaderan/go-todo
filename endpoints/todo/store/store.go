package store

import "todo/endpoints/todo/model"

var Todos = []model.Todo{
	{ID: 1, Title: "Learn Gin", Done: false},
	{ID: 2, Title: "Build Todo App", Done: false},
}

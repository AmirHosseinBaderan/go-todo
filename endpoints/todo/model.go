package todo

// Todo structure
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "First Todo", Done: false},
}

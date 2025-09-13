package todo

// Todo structure
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Done  bool   `json:"done"`
}

package data

type task struct {
	ID    string `json:"id"`
	Title string `json:"task"`
	Desc  string `json:"description"`
}

// Tasks - type Tasks
type Tasks []task

var toDoList = Tasks{
	{
		ID:    "1",
		Title: "My First Task",
		Desc:  "Only a task item for test",
	},
}

// GetTodoList - returns the TODO list
func GetTodoList() Tasks {
	return toDoList
}

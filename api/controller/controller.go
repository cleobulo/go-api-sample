package controller

import (
	"encoding/json"
	"go-api-sample/api/data"
	"net/http"
)

// GetTodoList - returns TODO list in JSON format
func GetTodoList(response http.ResponseWriter, request *http.Request) {
	todoList := data.GetTodoList()
	json.NewEncoder(response).Encode(todoList)
}

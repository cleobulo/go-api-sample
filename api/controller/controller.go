package controller

import (
	"encoding/json"
	"go-api-sample/api/data"
	"net/http"
)

// GetTest - implements only for test
func GetTest(response http.ResponseWriter, request *http.Request) {
	todoList := data.GetTodoList()
	json.NewEncoder(response).Encode(todoList)
}

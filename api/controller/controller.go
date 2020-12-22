package controller

import (
	"encoding/json"
	"net/http"
)

type task struct {
	ID    string `json:"id"`
	Title string `json:"task"`
	Desc  string `json:"description"`
}

type tasks []task

// GetTest - implements only for test
func GetTest(response http.ResponseWriter, request *http.Request) {
	allTasks := tasks{
		{
			ID:    "1",
			Title: "My First Task",
			Desc:  "Only a task item for test",
		},
	}
	json.NewEncoder(response).Encode(allTasks)
}

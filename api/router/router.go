package router

import (
	"go-api-sample/api/controller"

	"github.com/gorilla/mux"
)

// InitRoutes - Defines all the routes of the server
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todo", controller.GetTodoList).Methods("GET")
	return router
}

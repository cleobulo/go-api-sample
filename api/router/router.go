package router

import (
	"github.com/gorilla/mux"
)

// InitRoutes - Defines all the routes of the server
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	return router
}

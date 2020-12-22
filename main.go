package main

import (
	"fmt"
	"go-api-sample/core"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello! API works...")
	fmt.Printf("%s v%s\n", core.Application, core.Version)
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}

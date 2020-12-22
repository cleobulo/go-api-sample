package main

import (
	"fmt"
	"go-api-sample/api/router"
	"go-api-sample/core"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello! API works...")
	fmt.Printf("%s v%s\n", core.Application, core.Version)
	log.Fatal(http.ListenAndServe(core.ServerPort, router.InitRoutes()))
}

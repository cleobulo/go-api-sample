package main

import (
	"fmt"
	"go-api-sample/core"
)

func main() {
	fmt.Println("Hello! API works...")
	fmt.Printf("%s v%s\n", core.Application, core.Version)
}

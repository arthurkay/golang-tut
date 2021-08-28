package main

import (
	"os"

	"net/http"

	"db_proj/handlers"

	"github.com/arthurkay/env"
)

func main() {
	env.Load()
	port := os.Getenv("PORT")

	route := http.NewServeMux()

	route.HandleFunc("/", handlers.Home)
	route.HandleFunc("/new", handlers.New)
	route.HandleFunc("/first-user", handlers.First)
	err := http.ListenAndServe(":"+port, route)

	if err != nil {
		panic(err)
	}
}

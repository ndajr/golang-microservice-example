package main

import (
	"log"
	"net/http"

	server "github.com/neemiasjnr/gotest/pkg"
)

func main() {
	if err := http.ListenAndServe(":8080", server.New()); err != nil {
		log.Fatalln(err)
	}
}

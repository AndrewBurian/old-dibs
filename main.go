package main

import (
	"fmt"
	"net/http"
)

func main() {

	// static content
	http.Handle("/", http.FileServer(http.Dir("./views/")))

	// REST endpoints
	http.HandleFunc("/auth", authHandler)

	fmt.Println("Starting server!")
	http.ListenAndServe(":8080", nil)
}


package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Switch example with cases for Methods
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Connected to Server")
	case http.MethodPost:
		fmt.Fprintf(w, "Hello, POST method!")
	case http.MethodPut:
		fmt.Fprintf(w, "Hello, PUT method!")
	case http.MethodDelete:
		fmt.Fprintf(w, "Hello, DELETE method!")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Register the handler function for the root URL path
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/landing", helloHandler)

	PORT := ":8080"

	log.Printf("Starting server on %s\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

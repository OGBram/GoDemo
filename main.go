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
		fmt.Fprintf(w, "Connected to Server, GET method!")
	case http.MethodPost:
		fmt.Fprintf(w, "Landing, POST method!")
	case http.MethodPut:
		fmt.Fprintf(w, "Landing, PUT method!")
	case http.MethodDelete:
		fmt.Fprintf(w, "Landing, DELETE method!")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Register the frontend and handler function for the root URL path
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)
	// Register the handler function for the landing URL path
	http.HandleFunc("/landing", helloHandler)

	PORT := ":8080" //local

	log.Printf("Starting server on %s\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

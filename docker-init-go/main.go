package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the handler function for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	

	// Start the server on port 8080
	log.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
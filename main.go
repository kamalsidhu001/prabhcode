package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "templates" directory
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)

	// Additional handler for a custom route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, welcome to the server!")
	})

	// Define the port to run the server on
	port := ":80"
	fmt.Printf("Starting server on http://localhost%s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

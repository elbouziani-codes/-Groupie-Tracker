package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	// The HomeHandler handles the root path, while ArtistHandler manages detailed artist views.
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)

	// Custom Static Handler: Serves files from the "static" directory.
	http.HandleFunc("/static/", handlers.StaticHandlers)

	// Define the server address and port.
	fmt.Printf("Server successfully started at http://localhost:8080\n")

	// Initialize the HTTP server and listen for incoming requests.
	// log.Fatal is used to catch and report any critical startup errors.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Critical Error: Failed to start the server: %v", err)
	}
}

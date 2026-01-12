package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/handlers"
)

func main() {
	// Register application routes for the Home page and Artist details page.
	// The HomeHandler handles the root path, while ArtistHandler manages detailed artist views.
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)

	// Custom Static Handler: Serves files from the "static" directory.
	// This implementation ensures that missing files or directory access attempts
	// are handled by our custom ErrorHandler instead of the default Go 404 page.
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		// Construct the local file path by removing the leading slash.
		filePath := r.URL.Path[1:]

		// Verify if the requested file exists on the server.
		fileInfo, err := os.Stat(filePath)
		if os.IsNotExist(err) || fileInfo.IsDir() {
			// If the file is missing or a directory is requested, return a custom 404 error page.
			handlers.ErrorHandler(w, "Resource Not Found", http.StatusNotFound)
			return
		}

		// Serve the static file using the standard FileServer after stripping the "/static/" prefix.
		staticServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
		staticServer.ServeHTTP(w, r)
	})

	// Legacy support for direct CSS requests (optional if using /static/style.css in HTML).
	http.HandleFunc("/style.css", handlers.CssHandler)

	// Define the server address and port.
	serverAddr := ":8080"
	fmt.Printf("Server successfully started at http://localhost%s\n", serverAddr)

	// Initialize the HTTP server and listen for incoming requests.
	// log.Fatal is used to catch and report any critical startup errors.
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("Critical Error: Failed to start the server: %v", err)
	}
}

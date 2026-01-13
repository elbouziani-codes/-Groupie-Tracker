package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

// artists holds the list of artists (Note: using a global variable here might not be thread-safe for concurrent requests if modified elsewhere)
var artists []models.Artist

// HomeHandler handles requests to the root URL ("/")
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ensure the path is exactly "/", otherwise return 404
	if r.URL.Path != "/" {
		ErrorHandler(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// Parse the home page template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Fetch the list of artists from the API
	artists, err = utils.FetchArtists()
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template into a buffer to check for errors before writing to response
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, artists)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the buffered content to the response
	w.Write(buff.Bytes())
}

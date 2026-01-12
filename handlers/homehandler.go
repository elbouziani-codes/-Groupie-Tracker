package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

var artists []models.Artist

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	 if r.Method != http.MethodGet {
	 	ErrorHandler(w, "method not allowed", http.StatusMethodNotAllowed)
	 	return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, "404 Page not found", http.StatusNotFound)
	}
	// template.ParseFiles reades the html file when he found action like {{.}} he stocks in template object
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "internal server error", http.StatusInternalServerError)
		return
	}

	artists, err = utils.FetchArtists()
	if err != nil {
		ErrorHandler(w, "internal server error", http.StatusInternalServerError)
		return
	}
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, artists)
	if err != nil {
		ErrorHandler(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(buff.Bytes())
}

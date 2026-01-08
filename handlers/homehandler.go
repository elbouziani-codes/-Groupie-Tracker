package handlers

import (

	"html/template"
	"net/http"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

var artists []models.Artist

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	artists, err = utils.FetchArtists()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}



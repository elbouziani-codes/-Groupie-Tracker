package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/models"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	artistId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	found := false
	var selectedArtist models.Artist
	for _, v := range artists {
		if v.ID == artistId {
			selectedArtist = v
			found = true
			break
		}
	}
	if !(found) {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("template/artist.html")
	if err != nil {
		http.Error(w, "interanl server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		http.Error(w, "interanl server error", http.StatusInternalServerError)
		return
	}
}

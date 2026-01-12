package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		ErrorHandler(w, "Bad Request: Missing ID", http.StatusBadRequest)
		return
	}

	artistId, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, "Invalid Artist ID", http.StatusBadRequest)
		return
	}

	if len(artists) == 0 {
		artists, err = utils.FetchArtists()
		if err != nil {
			ErrorHandler(w, "Internal Server Error: Fetching Data", http.StatusInternalServerError)
			return
		}
	}

	var selectedArtist models.Artist
	found := false
	for _, v := range artists {
		if v.ID == artistId {
			selectedArtist = v
			found = true
			break
		}
	}

	if !found {
		ErrorHandler(w, "Artist Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error: Template Not Found", http.StatusInternalServerError)
		return
	}

	// Fetch relations data
	relations, err := utils.FetchRelations()
	if err != nil {
		fmt.Println("Error fetching relations:", err)
		ErrorHandler(w, "Internal Server Error: Fetching Relations", http.StatusInternalServerError)
		return
	}

	var datesLocations map[string][]string
	for _, rel := range relations.Index {
		if rel.ID == selectedArtist.ID {
			datesLocations = rel.DatesLocations
			fmt.Printf("Found relation for artist %d: %v\n", selectedArtist.ID, datesLocations)
			break
		}
	}
	if datesLocations == nil {
		fmt.Printf("No relation found for artist %d\n", selectedArtist.ID)
	}

	data := models.ArtistPageData{
		Artist:         selectedArtist,
		DatesLocations: datesLocations,
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, data)
	if err != nil {
		ErrorHandler(w, "Internal Server Error: Rendering Page", http.StatusInternalServerError)
		return
	}

	w.Write(buff.Bytes())
}

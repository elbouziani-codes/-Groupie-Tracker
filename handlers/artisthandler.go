package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

// ArtistHandler handles requests for specific artist details
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the "id" query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Validate that the id is a valid integer
	artistId, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check if the artist exists in the global 'artists' list
	found := false
	var selectedArtist models.Artist
	for _, v := range artists {
		if v.ID == artistId {
			selectedArtist = v
			found = true
			break
		}
	}

	// If the artist is not found, return a 404 error
	if !found {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	// Fetch location data associated with the artist
	location, err := utils.FetchLocations()
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	artistLocation := GetArtistLocation(location, artistId)

	// Fetch date data associated with the artist
	date, err := utils.FetchDates()
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	artistdate := GetArtistDate(date, artistId)

	// Fetch relation data associated with the artist
	artistrelation, err := utils.FetchRelations()
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	datarelation := GetArtistRelation(artistrelation, artistId)

	// Aggregate all data into a page data struct
	data := models.ArtistPageData{
		Artist:         selectedArtist,
		Location:       artistLocation,
		Date:           artistdate,
		DatesLocations: datarelation.DatesLocations,
	}

	// Parse the artist detail template
	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, data)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the result to the response
	w.Write(buff.Bytes())
}

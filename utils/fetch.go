package utils

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

// Base URLs for the API endpoints
const (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

// FetchData is a helper function to perform GET requests and decode JSON
func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// FetchArtists retrieves the full list of artists
func FetchArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := FetchData(ArtistsURL, &artists)
	return artists, err
}

// FetchRelations retrieves all concert relations
func FetchRelations() (models.RelationIndex, error) {
	var rels models.RelationIndex
	err := FetchData(RelationsURL, &rels)
	return rels, err
}

// FetchLocations retrieves all concert locations
func FetchLocations() (models.LocationIndex, error) {
	var locs models.LocationIndex
	err := FetchData(LocationsURL, &locs)
	return locs, err
}

// FetchDates retrieves all concert dates
func FetchDates() (models.DateIndex, error) {
	var dates models.DateIndex
	err := FetchData(DatesURL, &dates)
	return dates, err
}

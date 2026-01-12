package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/models"
	"groupie-tracker/utils"
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

	location, err := utils.FetchLocations()
	if err != nil {
		http.Error(w, "interanl server error", http.StatusInternalServerError)
		return
	}
	artistLocation:= GetArtistLocation(location, artistId)

	date, err := utils.FetchDates()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	artistdate:= GetArtistDate(date, artistId)
	artistrelation,err:=utils.FetchRelations()
	if err != nil{
		http.Error(w,"interanl server error",http.StatusInternalServerError)
	}
	datarelation:=GetArtidtRelation(artistrelation,artistId)

	data := models.ArtistPageData{
		Artist:   selectedArtist,
		Location: artistLocation,
		Date:     artistdate,
		DatesLocations: datarelation.DatesLocations,
	}

	tmpl, err := template.ParseFiles("template/artist.html")
	if err != nil {
		http.Error(w, "interanl server error", http.StatusInternalServerError)
		return
	}
	err =tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

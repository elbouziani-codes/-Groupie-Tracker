package handlers

import (
	"groupie-tracker/models"
)

func GetArtistLocation(locations models.LocationIndex, artistId int) (models.Location) {
	var locs models.Location
	for _, v := range locations.Index {
		if v.ID == artistId {
			locs = v
			return locs
		}
	}
	return locs
}
func GetArtistDate(dates models.DateIndex , artistId int) (models.Date){
	var Date models.Date
	for _, v := range dates.Index {
		if v.ID == artistId{
			Date = v
			return  Date 
		}
	}
	return Date
}
func GetArtidtRelation(relotion models.RelationIndex, artistId int) (models.Relation) {
	var datelocalition models.Relation
	for _,v := range relotion.Index{
		if v.ID == artistId{
			datelocalition = v
			return  datelocalition
		}
	}
	return datelocalition
}
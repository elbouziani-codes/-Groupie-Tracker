package handlers

import (
	"groupie-tracker/models"
)

// GetArtistLocation searches for an artist's location data by their ID
func GetArtistLocation(locations models.LocationIndex, artistId int) models.Location {
	var locs models.Location
	for _, v := range locations.Index {
		if v.ID == artistId {
			locs = v
			return locs
		}
	}
	return locs
}

// GetArtistDate searches for an artist's concert dates by their ID
func GetArtistDate(dates models.DateIndex, artistId int) models.Date {
	var date models.Date // Renamed from Date to date to avoid conflict with type name/convention
	for _, v := range dates.Index {
		if v.ID == artistId {
			date = v
			return date
		}
	}
	return date
}

// GetArtistRelation searches for an artist's relation data by their ID
// Renamed from GetArtidtRelation to fix typo
func GetArtistRelation(relations models.RelationIndex, artistId int) models.Relation {
	var relation models.Relation // Renamed from datelocalition for clarity
	for _, v := range relations.Index {
		if v.ID == artistId {
			relation = v
			return relation
		}
	}
	return relation
}

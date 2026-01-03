package models

// Artist represents the basic information of a band or artist
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Relation links the data of artists, dates, and locations
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

// Location represents concert locations
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationIndex struct {
	Index []Location `json:"index"`
}

// Date represents concert dates
type Date struct {
	
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DateIndex struct {
	Index []Date `json:"index"`
}

package dota

// BuildingStatus holds data about the endgame result of each teams buildings
type BuildingStatus struct {
	Barracks []Barrack `json:"barracks_status"`
	Towers   []Tower   `json:"tower_status"`
}

// A Barrack is... self-explanatory
type Barrack struct {
	Location  string `json:"location"`
	Destroyed bool   `json:"destroyed"`
}

// A Tower is a single tower
type Tower struct {
	Location  string `json:"location"`
	Destroyed bool   `json:"destroyed"`
}

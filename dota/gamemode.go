package dota

// GameMode is a specific type of Game - mapped to numbers by the WebAPI
type GameMode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GameModes is a list of all GameModes currently in DOTA 2
var GameModes = []GameMode{
	GameMode{
		ID:   0,
		Name: "No Game Mode",
	},
	GameMode{
		ID:   1,
		Name: "All Pick",
	},
	GameMode{
		ID:   2,
		Name: "Captain's Mode",
	},
	GameMode{
		ID:   3,
		Name: "Random Draft",
	},
	GameMode{
		ID:   4,
		Name: "Single Draft",
	},
	GameMode{
		ID:   5,
		Name: "All Random",
	},
	GameMode{
		ID:   7,
		Name: "Diretide",
	},
	GameMode{
		ID:   8,
		Name: "Reverse Captain's Mode",
	},
	GameMode{
		ID:   9,
		Name: "Greeviling",
	},
	GameMode{
		ID:   10,
		Name: "Tutorial",
	},
	GameMode{
		ID:   11,
		Name: "Mid Only",
	},
	GameMode{
		ID:   12,
		Name: "Least Played",
	},
	GameMode{
		ID:   13,
		Name: "New Player Pool",
	},
	GameMode{
		ID:   14,
		Name: "Compendium Matchmaking",
	},
	GameMode{
		ID:   15,
		Name: "Custom",
	},
	GameMode{
		ID:   16,
		Name: "Captain's Draft",
	},
	GameMode{
		ID:   17,
		Name: "Balanced Draft",
	},
	GameMode{
		ID:   18,
		Name: "Ability Draft",
	},
	GameMode{
		ID:   20,
		Name: "All Random Deathmatch",
	},
	GameMode{
		ID:   21,
		Name: "Solo Mid 1v1",
	},
	GameMode{
		ID:   22,
		Name: "Ranked All Pick",
	},
}

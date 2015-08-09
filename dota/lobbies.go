package dota

// A Lobby is what environment a match was played in, ie. Ranked or Unranked.
type Lobby struct {
	ID   int
	Name string
}

// Lobbies is a list of all lobby types currently in Dota 2
var Lobbies = []Lobby{
	Lobby{
		ID:   -1,
		Name: "Invalid",
	},
	Lobby{
		ID:   0,
		Name: "Public Matchmaking",
	},
	Lobby{
		ID:   1,
		Name: "Practice",
	},
	Lobby{
		ID:   2,
		Name: "Tournament",
	},
	Lobby{
		ID:   3,
		Name: "Tutorial",
	},
	Lobby{
		ID:   4,
		Name: "Co-Op with bots",
	},
	Lobby{
		ID:   5,
		Name: "Team Match",
	},
	Lobby{
		ID:   6,
		Name: "Solo Queue",
	},
	Lobby{
		ID:   7,
		Name: "Ranked",
	},
	Lobby{
		ID:   8,
		Name: "Solo Mid 1vs1",
	},
}

package dota

import "time"

// Match is a fully parsed dota 2 match
type Match struct {
	MatchID          int            `json:"match_id"`
	Dire             []Player       `json:"dire_players"`
	DireBuildings    BuildingStatus `json:"dire_building_status"`
	Radiant          []Player       `json:"radiant_players"`
	RadiantBuildings BuildingStatus `json:"radiant_building_status"`

	Victor     string    `json:"victor"`
	Duration   string    `json:"duration"`
	StartTime  time.Time `json:"start_time"`
	FirstBlood string    `json:"first_blood_time"`
	LobbyType  Lobby     `json:"lobby_type"`
	LeagueID   int       `json:"leagueid"`

	Draft *Draft `json:"draft,omitempty"`
	Lobby Lobby  `json:"lobby_type"`

	Region   Region   `json:"region"`
	GameMode GameMode `json:"game_mode"`
	Engine   string   `json:"game_engine"`
}

// Player is a player in a match
type Player struct {
	ID           int          `json:"player_id"`
	Hero         Hero         `json:"hero"`
	Kills        int          `json:"kills"`
	Deaths       int          `json:"deaths"`
	Assists      int          `json:"assists"`
	LeaverStatus LeaverStatus `json:"leaver_status"`
	Gold         int          `json:"gold"`
	LastHits     int          `json:"last_hits"`
	Denies       int          `json:"denies"`
	GPM          int          `json:"gpm"`
	XPM          int          `json:"xpm"`
	GoldSpent    int          `json:"gold_spent"`
	HeroDamage   int          `json:"hero_damage"`
	TowerDamage  int          `json:"tower_damage"`
	HeroHealing  int          `json:"hero_healing"`
	Level        int          `json:"level"`
	Items        []Item       `json:"items"`
	// in order, the skill build of this player this game
	SkillBuild []LevelUp   `json:"skill_build,omitempty"`
	SpiritBear *SpiritBear `json:"spirit_bear,omitempty"`
}

// The SpiritBear is the only additional unit currently in DOTA2 - it has
// it's own inventory, tied to a player, if that player is playing Lone Druid.
type SpiritBear struct {
	Items []Item `json:"items"`
}

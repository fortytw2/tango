package dota

// A Draft only exists for Captain's game modes
// for specific information on order, check out
// http://dota2.gamepedia.com/Game_modes#Captains_Mode
type Draft struct {
	FirstPick string `json:"first_pick"`

	RadiantPicks []int `json:"radiant_picks"`
	RadiantBans  []int `json:"radiant_bans"`

	DirePicks []int `json:"dire_picks"`
	DireBans  []int `json:"dire_bans"`
}

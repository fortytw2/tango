package request

// GetMatchHistory maps to the JSON returned from the API endpoint found at
// https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/V001/?key=<key>
// this endpoint returns very little data about the actual match
//
// the other MatchHistory endpoint uses the same JSON structure as this one ->
// http://api.steampowered.com/IDOTA2Match_<ID>/GetMatchHistoryBySequenceNum/V001/?key=<key>
// but has different parameters - not documented here, nor used by tango
/*
player_name=<name> # Search matches with a player name, exact match only
hero_id=<id> # Search for matches with a specific hero being played, hero id's are in dota/scripts/npc/npc_heroes.txt in your Dota install directory
skill=<skill>  # 0 for any, 1 for normal, 2 for high, 3 for very high skill
date_min=<date> # date in UTC seconds since Jan 1, 1970 (unix time format)
date_max=<date> # date in UTC seconds since Jan 1, 1970 (unix time format)
account_id=<id> # Steam account id (this is not SteamID, its only the account number portion)
league_id=<id> # matches for a particular league
start_at_match_id=<id> # Start the search at the indicated match id, descending
matches_requested=<n> # Defaults is 25 matches, this can limit to less
*/
type GetMatchHistory struct {
	Result struct {
		Status           int `json:"status"`
		NumResults       int `json:"num_results"`
		TotalResults     int `json:"total_results"`
		ResultsRemaining int `json:"results_remaining"`
		Matches          []struct {
			MatchID       int   `json:"match_id"`
			MatchSeqNum   int   `json:"match_seq_num"`
			StartTime     int64 `json:"start_time"`
			LobbyType     int   `json:"lobby_type"`
			RadiantTeamID int   `json:"radiant_team_id"`
			DireTeamID    int   `json:"dire_team_id"`
			Players       []struct {
				AccountID  int `json:"account_id"`
				PlayerSlot int `json:"player_slot"`
				HeroID     int `json:"hero_id"`
			} `json:"players"`
		} `json:"matches"`
	} `json:"result"`
}

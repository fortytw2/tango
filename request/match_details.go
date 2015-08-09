package request

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fortytw2/tango/dota"
)

// GetMatchDetails returns specific information about a single match, matching
// what's returned by https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/V001/?key=<key>&match_id=<match_id>
type GetMatchDetails struct {
	Result struct {
		Players []struct {
			// account_id - the player's 32-bit Steam ID - will be set to
			// "4294967295" if the account is private
			// or "0" if the player is a bot
			AccountID int `json:"account_id"`
			// player_slot - an 8-bit unsigned int: if the left-most bit is set,
			// the player was on dire. the two right-most bits represent the
			// player slot (0-4)
			PlayerSlot uint8 `json:"player_slot"`
			HeroID     int   `json:"hero_id"`
			// item_0 - the numeric ID of the item in the top-left slot.
			Item0 int `json:"item_0"`
			// item_1 - the numeric ID of the item in the top-center slot.
			Item1 int `json:"item_1"`
			// item_2 - the numeric ID of the item in the top-right slot.
			Item2 int `json:"item_2"`
			// item_3 - the numeric ID of the item in the bottom-left slot.
			Item3 int `json:"item_3"`
			// item_4 - the numeric ID of the item in the bottom-center slot.
			Item4 int `json:"item_4"`
			// item_5 - the numeric ID of the item in the bottom-right slot.
			Item5        int `json:"item_5"`
			Kills        int `json:"kills"`
			Deaths       int `json:"deaths"`
			Assists      int `json:"assists"`
			LeaverStatus int `json:"leaver_status"`
			Gold         int `json:"gold"`
			LastHits     int `json:"last_hits"`
			Denies       int `json:"denies"`
			GoldPerMin   int `json:"gold_per_min"`
			XpPerMin     int `json:"xp_per_min"`
			GoldSpent    int `json:"gold_spent"`
			HeroDamage   int `json:"hero_damage"`
			TowerDamage  int `json:"tower_damage"`
			HeroHealing  int `json:"hero_healing"`
			Level        int `json:"level"`
			// an array detailing the order in which a player's ability points were spent.
			AbilityUpgrades []struct {
				// the numeric ID of the ability that the point was spent on.
				Ability int `json:"ability"`
				// number of seconds since the start of the match.
				Time int `json:"time"`
				//  level of the hero when the ability was leveled.
				Level int `json:"level"`
			} `json:"ability_upgrades"`
			AdditionalUnits []struct {
				Unitname string `json:"unitname"`
				Item0    int    `json:"item_0"`
				Item1    int    `json:"item_1"`
				Item2    int    `json:"item_2"`
				Item3    int    `json:"item_3"`
				Item4    int    `json:"item_4"`
				Item5    int    `json:"item_5"`
			} `json:"additional_units"`
		} `json:"players"`
		Draft                 []Pick `json:"picks_bans"`
		RadiantWin            bool   `json:"radiant_win"`
		Duration              int    `json:"duration"`
		StartTime             int64  `json:"start_time"`
		MatchID               int    `json:"match_id"`
		MatchSeqNum           int    `json:"match_seq_num"`
		TowerStatusRadiant    uint16 `json:"tower_status_radiant"`
		TowerStatusDire       uint16 `json:"tower_status_dire"`
		BarracksStatusRadiant uint8  `json:"barracks_status_radiant"`
		BarracksStatusDire    uint8  `json:"barracks_status_dire"`
		Cluster               int    `json:"cluster"`
		FirstBloodTime        int    `json:"first_blood_time"`
		LobbyType             int    `json:"lobby_type"`
		HumanPlayers          int    `json:"human_players"`
		LeagueID              int    `json:"leagueid"`
		PositiveVotes         int    `json:"positive_votes"`
		NegativeVotes         int    `json:"negative_votes"`
		GameMode              int    `json:"game_mode"`
		Engine                int    `json:"engine"`
	} `json:"result"`
}

// A Pick is the way valve exposes picks/bans over the WebAPI.
type Pick struct {
	IsPick bool `json:"is_pick"`
	HeroID int  `json:"hero_id"`
	Team   int  `json:"team"`
	Order  int  `json:"order"`
}

func (md *GetMatchDetails) findPlayers() (radiant []dota.Player, dire []dota.Player, err error) {
	for _, p := range md.Result.Players {
		// find the proper hero...
		var hero dota.Hero
		for _, h := range dota.Heroes {
			if h.ID == p.HeroID {
				hero = h
				break
			}
		}

		if hero.ID == 0 {
			err = errors.New("no hero ID found")
			return
		}

		// find the proper items...
		var items []dota.Item
		for _, i := range dota.Items {
			playerItems := []int{p.Item0, p.Item1, p.Item2, p.Item3, p.Item4, p.Item5}
			for _, playerItem := range playerItems {
				if i.ID == playerItem {
					items = append(items, i)
				}
			}
		}

		if len(items) != 6 {
			err = fmt.Errorf("player %d does not have 6 items", p.AccountID)
			return
		}

		var skillBuild []dota.LevelUp
		for _, skill := range p.AbilityUpgrades {
			for _, ab := range dota.Abilities {
				if ab.ID == skill.Ability {
					// the numbers returned by the API are... not really right.
					// this attempts to account for the weirdness
					t, err := time.ParseDuration(strconv.Itoa(skill.Time-660) + "s")
					if err != nil {
						return nil, nil, err
					}

					sk := dota.LevelUp{
						ID:    ab.ID,
						Name:  ab.Name,
						Time:  t.String(),
						Level: skill.Level,
					}

					skillBuild = append(skillBuild, sk)
					break
				}
			}
		}

		player := dota.Player{
			ID:      p.AccountID,
			Hero:    hero,
			Kills:   p.Kills,
			Deaths:  p.Deaths,
			Assists: p.Assists,
			// LeaverStatus LeaverStatus
			Gold:        p.Gold,
			LastHits:    p.LastHits,
			Denies:      p.Denies,
			GPM:         p.GoldPerMin,
			XPM:         p.XpPerMin,
			GoldSpent:   p.GoldSpent,
			HeroDamage:  p.HeroDamage,
			TowerDamage: p.TowerDamage,
			HeroHealing: p.HeroHealing,
			Level:       p.Level,
			Items:       items,
			// in order, the skill build of this player this game
			SkillBuild: skillBuild,
			SpiritBear: nil,
		}

		if len(p.AdditionalUnits) > 0 {
			var spiritBear dota.SpiritBear

			var items []dota.Item
			for _, i := range dota.Items {
				playerItems := []int{p.AdditionalUnits[0].Item0, p.AdditionalUnits[0].Item1, p.AdditionalUnits[0].Item2, p.AdditionalUnits[0].Item3, p.AdditionalUnits[0].Item4, p.AdditionalUnits[0].Item5}
				for _, playerItem := range playerItems {
					if i.ID == playerItem {
						items = append(items, i)
					}
				}
			}

			if len(items) != 6 {
				err = fmt.Errorf("player %d does not have 6 items", p.AccountID)
				return
			}
			spiritBear.Items = items

			player.SpiritBear = &spiritBear
		}

		// let's just pretend these are ints...
		if p.PlayerSlot < 100 {
			radiant = append(radiant, player)
		} else {
			dire = append(dire, player)
		}
	}

	return
}

// getBuildingStatus parses Valve's format for tower and barracks status
// into something usable
// tower status is a 16 bit unsigned integer
/*
┌─┬─┬─┬─┬─────────────────────── Not used.
│ │ │ │ │ ┌───────────────────── Ancient Top
│ │ │ │ │ │ ┌─────────────────── Ancient Bottom
│ │ │ │ │ │ │ ┌───────────────── Bottom Tier 3
│ │ │ │ │ │ │ │ ┌─────────────── Bottom Tier 2
│ │ │ │ │ │ │ │ │ ┌───────────── Bottom Tier 1
│ │ │ │ │ │ │ │ │ │ ┌─────────── Middle Tier 3
│ │ │ │ │ │ │ │ │ │ │ ┌───────── Middle Tier 2
│ │ │ │ │ │ │ │ │ │ │ │ ┌─────── Middle Tier 1
│ │ │ │ │ │ │ │ │ │ │ │ │ ┌───── Top Tier 3
│ │ │ │ │ │ │ │ │ │ │ │ │ │ ┌─── Top Tier 2
│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ ┌─ Top Tier 1
│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
*/
// barracks status is an 8 bit unsigned integer
/*
┌─┬───────────── Not used.
│ │ ┌─────────── Bottom Ranged
│ │ │ ┌───────── Bottom Melee
│ │ │ │ ┌─────── Middle Ranged
│ │ │ │ │ ┌───── Middle Melee
│ │ │ │ │ │ ┌─── Top Ranged
│ │ │ │ │ │ │ ┌─ Top Melee
│ │ │ │ │ │ │ │
0 0 0 0 0 0 0 0
*/
func (md *GetMatchDetails) getBuildingStatus(dire bool) dota.BuildingStatus {
	bs := &dota.BuildingStatus{
		Towers:   make([]dota.Tower, 0),
		Barracks: make([]dota.Barrack, 0),
	}

	// binary love...
	var tower string
	var rax string
	if dire {
		tower = fmt.Sprintf("%b", md.Result.TowerStatusDire)
		rax = fmt.Sprintf("%b", md.Result.BarracksStatusDire)
	} else {
		tower = fmt.Sprintf("%b", md.Result.TowerStatusRadiant)
		rax = fmt.Sprintf("%b", md.Result.BarracksStatusRadiant)
	}

	for i, byt := range tower {
		switch i {
		case 0, 1, 2, 3, 4:
		case 5:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Top T4", Destroyed: isDestroyed(byt)})
		case 6:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Bot T4", Destroyed: isDestroyed(byt)})
		case 7:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Bot T3", Destroyed: isDestroyed(byt)})
		case 8:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Bot T2", Destroyed: isDestroyed(byt)})
		case 9:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Bot T1", Destroyed: isDestroyed(byt)})
		case 10:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Mid T3", Destroyed: isDestroyed(byt)})
		case 11:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Mid T2", Destroyed: isDestroyed(byt)})
		case 12:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Mid T1", Destroyed: isDestroyed(byt)})
		case 13:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Top T3", Destroyed: isDestroyed(byt)})
		case 14:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Top T2", Destroyed: isDestroyed(byt)})
		case 15:
			bs.Towers = append(bs.Towers, dota.Tower{Location: "Top T1", Destroyed: isDestroyed(byt)})
		}
	}

	for i, byt := range rax {
		switch i {
		case 0, 1:
		case 2:
			bs.Barracks = append(bs.Barracks, dota.Barrack{Location: "Bottom Ranged", Destroyed: isDestroyed(byt)})
		case 3:
			bs.Barracks = append(bs.Barracks, dota.Barrack{Location: "Bottom Melee", Destroyed: isDestroyed(byt)})
		case 4:
			bs.Barracks = append(bs.Barracks, dota.Barrack{Location: "Middle Ranged", Destroyed: isDestroyed(byt)})
		case 5:
			bs.Barracks = append(bs.Barracks, dota.Barrack{Location: "Middle Melee", Destroyed: isDestroyed(byt)})
		case 6:
			bs.Barracks = append(bs.Barracks, dota.Barrack{Location: "Top Ranged", Destroyed: isDestroyed(byt)})
		case 7:
			bs.Barracks = append(bs.Barracks, dota.Barrack{Location: "Top Melee", Destroyed: isDestroyed(byt)})
		}
	}
	return *bs
}

func isDestroyed(r rune) bool {
	if r == '1' {
		return true
	}
	return false
}

func (md *GetMatchDetails) parseDuration() (string, error) {
	t, err := time.ParseDuration(strconv.Itoa(md.Result.Duration) + "s")
	if err != nil {
		return "", err
	}
	return t.String(), nil
}

func (md *GetMatchDetails) parseFirstBloodTime() (string, error) {
	t, err := time.ParseDuration(strconv.Itoa(md.Result.FirstBloodTime) + "s")
	if err != nil {
		return "", err
	}
	return t.String(), nil
}

func (md *GetMatchDetails) findGameMode() (*dota.GameMode, error) {
	for _, gm := range dota.GameModes {
		if gm.ID == md.Result.GameMode {
			return &gm, nil
		}
	}
	return nil, fmt.Errorf("game mode ID %d not found", md.Result.GameMode)
}

func (md *GetMatchDetails) findRegion() (*dota.Region, error) {
	for _, r := range dota.Regions {
		if r.ID == md.Result.Cluster {
			return &r, nil
		}
	}
	return nil, fmt.Errorf("dota region ID %d not found", md.Result.Cluster)
}

func (md *GetMatchDetails) findLobbyType() (*dota.Lobby, error) {
	for _, l := range dota.Lobbies {
		if l.ID == md.Result.LobbyType {
			return &l, nil
		}
	}
	return nil, fmt.Errorf("lobby type ID %d not found", md.Result.LobbyType)
}

func (md *GetMatchDetails) getDraft() (*dota.Draft, error) {
	d := &dota.Draft{
		RadiantPicks: make([]int, 0),
		DirePicks:    make([]int, 0),
		RadiantBans:  make([]int, 0),
		DireBans:     make([]int, 0),
	}

	if len(md.Result.Draft) == 0 {
		return nil, nil
	}

	for _, pick := range md.Result.Draft {
		if pick.Order == 0 {
			switch pick.Team {
			case 0:
				d.FirstPick = "radiant"
			case 1:
				d.FirstPick = "dire"
			default:
				return nil, fmt.Errorf("unknown team %d", pick.Team)
			}
		}

		if pick.IsPick {
			switch pick.Team {
			case 0:
				d.RadiantPicks = append(d.RadiantPicks, pick.HeroID)
			case 1:
				d.DirePicks = append(d.DirePicks, pick.HeroID)
			default:
				return nil, fmt.Errorf("unknown team %d", pick.Team)
			}
		} else {
			switch pick.Team {
			case 0:
				d.RadiantBans = append(d.RadiantBans, pick.HeroID)
			case 1:
				d.DireBans = append(d.DireBans, pick.HeroID)
			default:
				return nil, fmt.Errorf("unknown team %d", pick.Team)
			}
		}
	}
	return d, nil
}

// ParseMatch turns the Valve JSON matchdetails into a usable *dota.Match
func (md *GetMatchDetails) ParseMatch() (*dota.Match, error) {
	var victor string
	if md.Result.RadiantWin {
		victor = "radiant"
	} else {
		victor = "dire"
	}

	var engine string
	if md.Result.Engine == 0 {
		engine = "Source 1"
	} else if md.Result.Engine == 1 {
		engine = "Source 2"
	} else {
		return nil, fmt.Errorf("unknown engine %d", md.Result.Engine)
	}

	radiantPlayers, direPlayers, err := md.findPlayers()
	if err != nil {
		return nil, err
	}

	gm, err := md.findGameMode()
	if err != nil {
		return nil, err
	}

	region, err := md.findRegion()
	if err != nil {
		return nil, err
	}

	l, err := md.findLobbyType()
	if err != nil {
		return nil, err
	}

	dur, err := md.parseDuration()
	if err != nil {
		return nil, err
	}

	fbt, err := md.parseFirstBloodTime()
	if err != nil {
		return nil, err
	}

	draft, err := md.getDraft()
	if err != nil {
		return nil, err
	}

	match := &dota.Match{
		MatchID:          md.Result.MatchID,
		Dire:             direPlayers,
		DireBuildings:    md.getBuildingStatus(true),
		Radiant:          radiantPlayers,
		RadiantBuildings: md.getBuildingStatus(false),
		Victor:           victor,
		Draft:            draft,
		StartTime:        time.Unix(md.Result.StartTime, 0),
		Duration:         dur,
		GameMode:         *gm,
		Lobby:            *l,
		FirstBlood:       fbt,
		Engine:           engine,
		Region:           *region,
		LeagueID:         md.Result.LeagueID,
	}

	return match, nil
}

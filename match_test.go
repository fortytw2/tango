package tango

import (
	"log"
	"os"
	"testing"
	"time"
)

// matches covers all matches we're making test cases for
var matches = []int{
	// TI3 Grand Finals, all 5 games (ordered 1-5)
	271076032,
	271088718,
	271102834,
	271123757,
	271145478,
	// TI4 Grand Finals, all 4 games (ordered 1-4)
	789453600,
	789518247,
	789590883,
	789645621,
	// DAC Grand Finals (3 games)
	1223581106,
	1223686866,
	1223796316,
}

func TestGetMatch(t *testing.T) {
	if os.Getenv("STEAM_API_KEY") == "" {
		t.Fatal("no steam API key found, cannot test")
	}

	api := NewWebAPI(os.Getenv("STEAM_API_KEY"))

	// ensure there are no errors getting matches
	for _, id := range matches {
		_, err := api.GetMatch(id)
		if err != nil {
			t.Errorf("error getting match %d, %s", id, err)
		}
	}
}

func TestGetPlayerMatches(t *testing.T) {
	if os.Getenv("STEAM_API_KEY") == "" {
		t.Fatal("no steam API key found, cannot test")
	}

	api := NewWebAPI(os.Getenv("STEAM_API_KEY"))

	ids, err := api.GetPlayerMatchIDs(60746141, time.Date(2015, time.May, 10, 23, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("error getting match IDs for player %d, %s", 60746141, err)
	}

	log.Println(len(ids))
}

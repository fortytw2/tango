package tango

import (
	"os"
	"testing"
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

	for _, id := range matches {
		m, err := api.GetMatch(id)
		if err != nil {
			t.Errorf("error getting match %d, %s", id, err)
		}
	}
}

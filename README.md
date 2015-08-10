Tango [![GoDoc](https://godoc.org/github.com/fortytw2/tango?status.svg)](http://godoc.org/github.com/fortytw2/tango)
------

Complete wrapper for Valve's DOTA 2 Web API exposed cleanly to Go.

Supports all available WebAPI methods except for `GetTournamentPlayerStats` and
`GetTournamentPrizePool` as they both only work for `LeagueID=65006`, which is
The International 2013.

### Example

```go
package main

import (
    "fmt"
	"log"

	"github.com/fortytw2/tango"
)

var matchID = 271145478

func main() {
	api := tango.NewWebAPI("API KEY")

	m, err := api.GetMatch(matchID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DOTA 2 Match", matchID)
	fmt.Println("Winner -", m.Victor)

	fmt.Println("Radiant Heroes")
	for _, player := range m.Radiant {
		fmt.Println(player.Hero.LocalizedName)
	}

	fmt.Println("Dire Heroes")
	for _, player := range m.Dire {
		fmt.Println(player.Hero.LocalizedName)
	}
}
```

Should be very straightforward to use, check Godoc for complete documentation
and open an issue here on GitHub if anything goes wrong.

LICENSE
------
MIT, see LICENSE

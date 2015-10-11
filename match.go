package tango

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/fortytw2/tango/dota"
	"github.com/fortytw2/tango/request"
)

// GetMatch returns a *dota.Match parsed out of the Dota2 WebAPI.
func (wa *WebAPI) GetMatch(id int) (*dota.Match, error) {
	url := &url.URL{
		Scheme: "https",
		Path:   "api.steampowered.com/IDOTA2Match_570/GetMatchDetails/V001/",
	}
	query := url.Query()
	query.Set("key", wa.key)
	query.Set("match_id", strconv.Itoa(id))
	url.RawQuery = query.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var md request.GetMatchDetails
	err = json.Unmarshal(data, &md)
	if err != nil {
		return nil, err
	}

	return md.ParseMatch()
}

// GetPlayerMatchIDs gets all of a players matchIDs since a certain time,
// because this requires potentially a large number of HTTP requests,
// the function will block for a while.
// Unfortunately, Valve currently restricts the GetMatchHistory endpoint to
// a mere 500 entries, returning the last match repeated if you try to go
// further. It appears there are solutions for this, just need to implement
// them here - this code isn't the greatest anyways...
func (wa *WebAPI) GetPlayerMatchIDs(playerID int, since time.Time) ([]int, error) {
	url := &url.URL{
		Scheme: "https",
		Path:   "api.steampowered.com/IDOTA2Match_570/GetMatchHistory/V001/",
	}
	query := url.Query()
	query.Set("key", wa.key)
	query.Set("account_id", strconv.Itoa(playerID))
	url.RawQuery = query.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mh request.GetMatchHistory
	err = json.Unmarshal(data, &mh)
	if err != nil {
		return nil, err
	}

	var ids []int
	for _, m := range mh.Result.Matches {
		if time.Unix(m.StartTime, 0).Before(since) {
			return ids, nil
		}
		ids = append(ids, m.MatchID)
	}
	// continue getting matches until we've reached the end, or since
	for {
		query := url.Query()
		query.Set("key", wa.key)
		query.Set("account_id", strconv.Itoa(playerID))
		query.Set("start_at_match_id", strconv.Itoa(ids[len(ids)-1]))
		url.RawQuery = query.Encode()

		resp, err := http.Get(url.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var mh request.GetMatchHistory
		err = json.Unmarshal(data, &mh)
		if err != nil {
			return nil, err
		}

		if mh.Result.ResultsRemaining == 0 {
			return ids, nil
		}

		for _, m := range mh.Result.Matches {
			if time.Unix(m.StartTime, 0).Before(since) {
				return ids, nil
			}
			ids = append(ids, m.MatchID)
		}

		log.Println(len(ids))
	}
}

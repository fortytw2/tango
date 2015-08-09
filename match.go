package tango

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

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

package request

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestParseMatchDetails(t *testing.T) {
	files, err := filepath.Glob("../fixtures/*.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}

		var md GetMatchDetails
		err = json.Unmarshal(data, &md)
		if err != nil {
			t.Fatal(err)
		}

		_, err = md.ParseMatch()
		if err != nil {
			t.Fatal("could not parse match!", file, err)
		}
	}

	return
}

package tango

// WebAPI just wraps the parsing functionality found in requests/
// in an easy to use API
type WebAPI struct {
	key string
}

// NewWebAPI creates a new instance of WebAPI, with the given steam development
// key held internally. http://steamcommunity.com/dev/registerkey
func NewWebAPI(key string) *WebAPI {
	return &WebAPI{
		key: key,
	}
}

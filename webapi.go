package tango

type WebAPI struct {
	key string
}

func NewWebAPI(key string) *WebAPI {
	return &WebAPI{
		key: key,
	}
}

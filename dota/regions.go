package dota

// A Region is a cluster location
// more than likely, IP addresses can be added in here
type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Regions is a list of all current Dota2 Regions
var Regions = []Region{
	Region{
		ID:   111,
		Name: "US West",
	},
	Region{
		ID:   112,
		Name: "US West",
	},
	Region{
		ID:   114,
		Name: "US West",
	},
	Region{
		ID:   121,
		Name: "US East",
	},
	Region{
		ID:   122,
		Name: "US East",
	},
	Region{
		ID:   123,
		Name: "US East",
	},
	Region{
		ID:   124,
		Name: "US East",
	},
	Region{
		ID:   131,
		Name: "Europe West",
	},
	Region{
		ID:   132,
		Name: "Europe West",
	},
	Region{
		ID:   133,
		Name: "Europe West",
	},
	Region{
		ID:   134,
		Name: "Europe West",
	},
	Region{
		ID:   135,
		Name: "Europe West",
	},
	Region{
		ID:   136,
		Name: "Europe West",
	},
	Region{
		ID:   142,
		Name: "South Korea",
	},
	Region{
		ID:   143,
		Name: "South Korea",
	},
	Region{
		ID:   151,
		Name: "Southeast Asia",
	},
	Region{
		ID:   152,
		Name: "Southeast Asia",
	},
	Region{
		ID:   153,
		Name: "Southeast Asia",
	},
	Region{
		ID:   161,
		Name: "China",
	},
	Region{
		ID:   163,
		Name: "China",
	},
	Region{
		ID:   171,
		Name: "Australia",
	},
	Region{
		ID:   181,
		Name: "Russia",
	},
	Region{
		ID:   182,
		Name: "Russia",
	},
	Region{
		ID:   183,
		Name: "Russia",
	},
	Region{
		ID:   184,
		Name: "Russia",
	},
	Region{
		ID:   185,
		Name: "Russia",
	},
	Region{
		ID:   186,
		Name: "Russia",
	},
	Region{
		ID:   191,
		Name: "Europe East",
	},
	Region{
		ID:   192,
		Name: "Europe East",
	},
	Region{
		ID:   200,
		Name: "South America",
	},
	Region{
		ID:   202,
		Name: "South America",
	},
	Region{
		ID:   203,
		Name: "South America",
	},
	Region{
		ID:   204,
		Name: "South America",
	},
	Region{
		ID:   211,
		Name: "South Africa",
	},
	Region{
		ID:   212,
		Name: "South Africa",
	},
	Region{
		ID:   213,
		Name: "South Africa",
	},
	Region{
		ID:   221,
		Name: "China",
	},
	Region{
		ID:   222,
		Name: "China",
	},
	Region{
		ID:   223,
		Name: "China",
	},
	Region{
		ID:   224,
		Name: "China",
	},
	Region{
		ID:   225,
		Name: "China",
	},
	Region{
		ID:   231,
		Name: "China",
	},
	Region{
		ID:   242,
		Name: "Chile",
	},
	Region{
		ID:   251,
		Name: "Peru",
	},
	Region{
		ID:   261,
		Name: "India",
	},
}

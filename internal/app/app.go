package app

type PokedexEntry struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type pokedex map[string]PokedexEntry

type config struct {
	NextLocationUrl     string
	PreviousLocationUrl string
	Pokedex             pokedex
}

var Config config

func Init() {
	pokedex := make(pokedex)
	Config = config{
		NextLocationUrl:     "https://pokeapi.co/api/v2/location-area/",
		PreviousLocationUrl: "",
		Pokedex:             pokedex,
	}
}

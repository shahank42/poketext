package pokeapi

import (
	"encoding/json"
)

type MapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string) (MapData, error) {
	data, _, err := getWithCache(url)
	if err != nil {
		return MapData{}, err
	}

	mapData := MapData{}
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		return MapData{}, err
	}

	return mapData, nil
}

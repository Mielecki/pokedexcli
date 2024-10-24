package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiURL = "https://pokeapi.co/api/v2"
)

type LocationAreaBody struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(locationURL *string) (LocationAreaBody, error) {
	url := apiURL + "/location-area"
	if locationURL != nil {
		url = *locationURL
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaBody{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaBody{}, err
	}
	
	locations := LocationAreaBody{}

	if err := json.Unmarshal(data, &locations); err != nil {
		fmt.Println("unmarshal")
		return LocationAreaBody{}, err
	}

	return locations, nil

}
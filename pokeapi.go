package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationResponse struct {
	Count    int     `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

func fetchLocation(url string) (LocationResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	var data LocationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationResponse{}, err
	}
	return data, nil
}

type LocationArea struct {
	GameIndex int    `json:"game_index,omitempty"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Location  struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters,omitempty"`
}

func fetchAreaPokemons(area string) (*LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("area %s not found", area)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var locationArea LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return nil, err
	}

	return &locationArea, nil
}

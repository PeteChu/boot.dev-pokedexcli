package main

import (
	"encoding/json"
	"fmt"
)

func (app *App) commandExit(args ...string) error {
	return ErrExit
}

func (app *App) commandHelp(args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Print all commands dynamically
	for _, cmd := range app.Commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func (app *App) commandMap(args ...string) error {
	url := "https://pokeapi.co/api/v2/location"
	if app.Locations.Next != nil {
		url = *app.Locations.Next
	}

	var locations LocationResponse

	// return from cache if exists
	if locs, exist := app.Cache.Get(url); exist {
		if err := json.Unmarshal(locs, &locations); err != nil {
			return err
		}
		app.Locations = locations
	} else {
		// fetch from api
		locations, err := fetchLocation(url)
		if err != nil {
			return err
		}
		data, err := json.Marshal(locations)
		if err != nil {
			return err
		}
		app.Cache.Add(url, data)
		app.Locations = locations
	}

	for _, l := range app.Locations.Results {
		fmt.Printf("%s-area\n", l.Name)
	}
	return nil
}

func (app *App) commandMapb(args ...string) error {
	url := "https://pokeapi.co/api/v2/location"
	if app.Locations.Previous != nil {
		url = *app.Locations.Previous
	}

	var locations LocationResponse

	// return from cache if exists
	if locs, exist := app.Cache.Get(url); exist {
		if err := json.Unmarshal(locs, &locations); err != nil {
			return err
		}
		app.Locations = locations
	} else {
		// fetch from api
		locations, err := fetchLocation(url)
		if err != nil {
			return err
		}
		data, err := json.Marshal(locations)
		if err != nil {
			return err
		}
		app.Cache.Add(url, data)
		app.Locations = locations
	}

	for _, l := range app.Locations.Results {
		fmt.Printf("%s-area\n", l.Name)
	}
	return nil
}

func (app *App) explore(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("area must be provide to explore")
	}

	area := args[0]
	loc, err := fetchAreaPokemons(area)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, pokemon := range loc.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

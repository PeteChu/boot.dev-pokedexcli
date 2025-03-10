package cli

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"pokedexcli/internal/api/pokeapi"
)

// CommandExit handles the exit command
func CommandExit(app *App, args ...string) error {
	return ErrExit
}

// CommandHelp displays available commands
func CommandHelp(app *App, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Print all commands dynamically
	for _, cmd := range app.Commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

// CommandMap displays the next page of locations
func CommandMap(app *App, args ...string) error {
	url := "https://pokeapi.co/api/v2/location"
	if app.Locations.Next != nil {
		url = *app.Locations.Next
	}

	var locations pokeapi.LocationResponse

	// return from cache if exists
	if locs, exist := app.Cache.Get(url); exist {
		if err := json.Unmarshal(locs, &locations); err != nil {
			return err
		}
		app.Locations = locations
	} else {
		// fetch from api
		var err error
		locations, err = app.Client.GetLocationArea(url)
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

// CommandMapb displays the previous page of locations
func CommandMapb(app *App, args ...string) error {
	url := "https://pokeapi.co/api/v2/location"
	if app.Locations.Previous != nil {
		url = *app.Locations.Previous
	}

	var locations pokeapi.LocationResponse

	// return from cache if exists
	if locs, exist := app.Cache.Get(url); exist {
		if err := json.Unmarshal(locs, &locations); err != nil {
			return err
		}
		app.Locations = locations
	} else {
		// fetch from api
		var err error
		locations, err = app.Client.GetLocationArea(url)
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

// CommandExplore lists Pokémon in a specific area
func CommandExplore(app *App, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("area must be provided to explore")
	}

	area := args[0]
	loc, err := app.Client.ExploreArea(area)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, pokemon := range loc.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func CommandCatch(app *App, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("what pokemon you want to catch?")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := app.Client.GetPokemonDetail(name)
	if err != nil {
		return err
	}

	// Calculate catch threshold - normalize it to ensure it's between 0 and 90
	// This ensures even high base experience Pokémon can be caught
	threshold := math.Min(90, float64(pokemon.BaseExperience)/3)

	// Random attempt value between 0 and 100
	catchAttempt := rand.Intn(100)

	fmt.Printf("Attempting to catch... (Difficulty: %.1f, Attempt: %d)\n", threshold, catchAttempt)

	// Determine if catch was successful
	if float64(catchAttempt) >= threshold {
		fmt.Printf("Caught %s!\n", name)
		// Add pokemon to collection
		app.Pokedex[pokemon.Name] = Pokemon{pokemon.Name}
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

var ErrExit = errors.New("Closing the Pokedex... Goodbye!")

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type LocationResponse struct {
	Count    int     `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

type App struct {
	Locations LocationResponse
	Commands  map[string]cliCommand
	Cache     *pokecache.Cache
}

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	defer cache.Stop()

	app := App{Cache: cache}

	command := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    app.commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    app.commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location areas",
			callback:    app.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous location areas",
			callback:    app.commandMapb,
		},
	}

	app.Commands = command

	for {
		s := bufio.NewScanner(os.Stdin)

		// Read first command
		fmt.Print("Pokedex > ")
		for s.Scan() {
			texts := cleanInput(s.Text())

			if len(texts) == 0 {
				fmt.Print("Pokedex > ")
				continue
			}

			command, ok := app.Commands[texts[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.callback()
			switch err {
			case ErrExit:
				fmt.Println(err.Error())
				os.Exit(0)
			}

			// Wait for next command
			fmt.Print("Pokedex > ")
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(
		strings.ToLower(text),
	)
}

func (app *App) commandExit() error {
	return ErrExit
}

func (app *App) commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Print all commands dynamically
	for _, cmd := range app.Commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func (app *App) commandMap() error {
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

func (app *App) commandMapb() error {
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

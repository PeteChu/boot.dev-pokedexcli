package cli

import (
	"errors"
	"pokedexcli/internal/api/pokeapi"
	"pokedexcli/internal/pokecache"
)

// ErrExit is returned when the user wants to exit the application
var ErrExit = errors.New("closing the Pokedex... goodbye")

// Command represents a CLI command with its metadata and handler
type Command struct {
	Name        string
	Description string
	Callback    func(app *App, args ...string) error
}

type Pokemon struct {
	Name   string         `json:"name"`
	Height int            `json:"height"`
	Weight int            `json:"weight"`
	Stats  []pokeapi.Stat `json:"stats"`
	Types  []pokeapi.Type `json:"types"`
}

// App holds the application state
type App struct {
	Client    *pokeapi.Client
	Cache     *pokecache.Cache
	Locations pokeapi.LocationResponse
	Commands  map[string]Command
	Pokedex   map[string]Pokemon
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

var ErrExit = errors.New("closing the Pokedex... goodbye")

type cliCommand struct {
	name        string
	description string
	callback    func(args ...string) error
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
		"explore": {
			name:        "explore",
			description: "List all the Pokémon located in the provided area",
			callback:    app.explore,
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
			args := texts[1:]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := command.callback(args...)
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

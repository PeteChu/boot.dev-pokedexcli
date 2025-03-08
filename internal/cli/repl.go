package cli

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/api/pokeapi"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

// StartRepl initializes and runs the REPL (Read-Eval-Print Loop)
func StartRepl() {
	cache := pokecache.NewCache(5 * time.Minute)
	defer cache.Stop()

	app := &App{
		Client: pokeapi.NewClient(),
		Cache:  cache,
	}

	app.Commands = map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays location areas",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays previous location areas",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "List all the Pokémon located in the provided area",
			Callback:    CommandExplore,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := input[1:]

		command, exists := app.Commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback(app, args...)
		if err == ErrExit {
			fmt.Println(err.Error())
			os.Exit(0)
		} else if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

// cleanInput normalizes and splits user input
func cleanInput(text string) []string {
	return strings.Fields(
		strings.ToLower(text),
	)
}

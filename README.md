# Pokedex CLI

Pokedex CLI is a command-line tool written in Go that lets you explore the world of Pokémon. Using data from the public PokéAPI (https://pokeapi.co/), you can browse location areas, explore Pokémon encounters, catch and inspect Pokémon, and maintain your own Pokedex. The application also features a built-in cache to improve performance when querying the API.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation & Setup](#installation--setup)
- [Usage](#usage)
  - [Available Commands](#available-commands)
- [Testing](#testing)
- [Repository Structure](#repository-structure)
- [Notes](#notes)
- [License](#license)

---

## Overview

Pokedex CLI is designed as an interactive terminal application. Users can issue commands at the prompt to perform various actions:
- **Map exploration:** Browse available location areas (next and previous pages).
- **Area exploration:** Explore a specific area to see which Pokémon can be encountered.
- **Catching Pokémon:** Attempt to catch Pokémon using a randomized threshold mechanism.
- **Inspecting Pokémon:** Get details like height, weight, stats, and types of caught Pokémon.
- **Viewing your Pokedex:** List all Pokémon you have caught.

The package is organized to separate different concerns:
- The `cmd/pokedex` folder contains the entry point.
- The `internal/api/pokeapi` package handles communication with the Pokémon API.
- The `internal/cli` package manages the command-line parsing and command callbacks.
- The `internal/pokecache` package implements a simple caching mechanism with expiration.
- Unit tests are provided in the `tests` directory.

---

## Features

- **Interactive REPL:** A command prompt for real-time command execution.
- **API Integration:** Communicate with the PokéAPI for up-to-date Pokémon data.
- **Caching:** Cache API responses to minimize redundant network calls.
- **Command-Based Navigation:** Easily list, explore, catch, inspect, and display Pokémon in your collection.
- **Extensible Commands:** The code structure allows for adding new commands easily.

---

## Installation & Setup

### Prerequisites

- Go 1.24.0 or newer must be installed on your machine. You can download it from [golang.org](https://golang.org/).

### Clone the Repository

```bash
git clone https://github.com/yourusername/pokedexcli.git
cd pokedexcli
```

### Build and Run

To run the application directly using Go:

```bash
go run cmd/pokedex/main.go
```

Alternatively, build the binary:

```bash
go build -o pokedexcli cmd/pokedex/main.go
./pokedexcli
```

---

## Usage

Upon launching the application, you will be greeted with a prompt:

```
Pokedex >
```

You can then type commands to interact with the application. The tool supports various commands that allow you to explore locations, catch Pokémon, inspect details, and manage your Pokedex.

### Available Commands

- **help:**  
  Displays available commands and a brief description of each.
  
  Example:
  ```
  Pokedex > help
  ```

- **map:**  
  Displays the current (or next) set of location areas from the PokéAPI.
  
  Example:
  ```
  Pokedex > map
  ```

- **mapb:**  
  Displays the previous set of location areas.
  
  Example:
  ```
  Pokedex > mapb
  ```

- **explore [area]:**  
  Shows all the Pokémon available in the specified location area.
  
  Example:
  ```
  Pokedex > explore kanto
  ```

- **catch [pokemon_name]:**  
  Attempts to catch a specified Pokémon. This uses a randomized threshold based on base experience.
  
  Example:
  ```
  Pokedex > catch pikachu
  ```

- **inspect [pokemon_name]:**  
  Displays detailed information about a caught Pokémon.
  
  Example:
  ```
  Pokedex > inspect pikachu
  ```

- **pokedex:**  
  Lists Pokémon that you have successfully caught.
  
  Example:
  ```
  Pokedex > pokedex
  ```

- **exit:**  
  Exits the application.
  
  Example:
  ```
  Pokedex > exit
  ```

---

## Testing

The repository includes tests for key functionality such as caching and input cleaning.

Run tests with:

```bash
go test ./tests/...
```

Make sure you have all dependencies installed. The tests cover:
- Cache operations (Add, Get, and expiration).
- Input cleaning function for the REPL.

---

## Repository Structure

- **cmd/pokedex/main.go:**  
  The entry point for the Pokedex CLI application.

- **internal/api/pokeapi:**  
  Contains API client implementations and data models for interacting with the PokéAPI.

  - `client.go` – Implements HTTP requests to fetch locations, areas, and Pokémon details.
  - `models.go` – Data models for API responses.

- **internal/cli:**  
  Manages command processing and the REPL interface.
  
  - `commands.go` – Implementation of action callbacks for each command.
  - `repl.go` – Initializes and manages the REPL loop.
  - `types.go` – Contains shared types such as the command structure, application state, and error definitions.

- **internal/pokecache/pokecache.go:**  
  Provides caching functionality with automatic expiration.

- **tests:**  
  Automated tests for validating cache functionality and REPL input processing.

- **.gitignore & .repomixignore:**  
  Configuration files for ignoring certain files from version control and bundled packaging.

- **go.mod:**  
  Go module definition for dependency management.

---

## Notes

- Some files (such as build and IDE-specific artefacts) are excluded from this packed representation.
- The repository respects common ignore patterns to prevent unnecessary files from being included in the build.
- Comments and empty lines have been removed in this representation to provide a concise overview of the codebase.

---

## License

Distributed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

Enjoy your adventure into the world of Pokémon with Pokedex CLI!

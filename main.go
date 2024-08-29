package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/kekekekyle/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand{
  return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
    "map": {
        name:        "map",
        description: "Returns the next 20 locations",
        callback:    commandMap,
    },
    "mapb": {
        name:        "mapb",
        description: "Returns the previous 20 locations",
        callback:    commandMapB,
    },
    "explore": {
        name:        "explore",
        description: "Explore the given location",
        callback:    commandExplore,
    },
    "catch": {
        name:        "catch",
        description: "Throw a pokeball at the given pokemon",
        callback:    commandCatch,
    },
    "inspect": {
        name:        "inspect",
        description: "Inspects the given pokemon",
        callback:    commandInspect,
    },
    "pokedex": {
        name:        "pokedex",
        description: "Lists all the pokemon in the pokedex",
        callback:    commandPokedex,
    },
  }
}

type Config struct {
  pokeApi pokeapi.Client
  NextURL *string
  PrevURL *string
  pokedex map[string]pokeapi.PokemonResponse
}

func start() {
  scanner := bufio.NewScanner(os.Stdin)
  pokeClient := pokeapi.NewClient(5 * time.Second)
  config := &Config{
    pokeApi: pokeClient,
    pokedex: map[string]pokeapi.PokemonResponse{},
  }
  commands := getCommands()

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    inputs := scanner.Text()
    input := strings.Split(inputs, " ")
    inputCommand := input[0] 
    arg1 := ""
    if len(input) > 1 {
      arg1 = input[1]
    }
    if command, exists := commands[inputCommand]; exists {
      err := command.callback(config, arg1)
      if err != nil {
        fmt.Println("Error!")
        break
      }
    } else {
      fmt.Println("Invalid command.")
    }
  }
}

func main() {
  start()
}

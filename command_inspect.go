package main

import (
  "fmt"
  "strconv"
)

func commandInspect(cfg *Config, args... string) error {
  if len(args) == 0 {
    return nil
  }
  inspect := args[0]
  if pokemon, ok := cfg.pokedex[inspect]; ok {
    fmt.Println("Name: " + pokemon.Name)
    fmt.Println("Height: " + strconv.Itoa(pokemon.Height))
    fmt.Println("Weight: " + strconv.Itoa(pokemon.Weight))
    fmt.Println("Stats: ")
    for _, stat := range pokemon.Stats {
      fmt.Println("- " + stat.Stat.Name + ": " + strconv.Itoa(stat.BaseStat))
    }
    fmt.Println("Types: ")
    for _, pokeType := range pokemon.Types {
      fmt.Println("- " + pokeType.Type.Name)
    }
  }

  return nil
}


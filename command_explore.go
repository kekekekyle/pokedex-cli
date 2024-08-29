package main

import (
  "fmt"
)

func commandExplore(cfg *Config, args... string) error {
  if len(args) == 0 {
    return nil
  }
  response, err := cfg.pokeApi.GetLocationName(args[0])
  if err != nil {
    return nil
  }

  fmt.Println("Pokemon in the area: ")
  for _, pokemon := range response.PokemonEncounters {
    fmt.Println("- " + pokemon.Pokemon.Name)
  }

  return nil
}


package main

import (
  "fmt"
  "math/rand"
)

func commandCatch(cfg *Config, args... string) error {
  if len(args) == 0 {
    return nil
  }
  pokemon := args[0]
  response, err := cfg.pokeApi.GetPokemon(pokemon)
  if err != nil {
    return nil
  }

  baseExperience := response.BaseExperience
  catchRate := rand.Intn(baseExperience)

  fmt.Println("Throwing a Pokeball at " + pokemon +  "...")
  if catchRate > baseExperience / 2 {
    fmt.Println(pokemon + " was caught!")
    cfg.pokedex[pokemon] = response
  } else {
    fmt.Println(pokemon + " escaped!")
  }

  return nil
}


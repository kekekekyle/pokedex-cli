package main

import (
  "fmt"
)

func commandMap(cfg *Config, args... string) error {
  response, err := cfg.pokeApi.ListLocations(cfg.NextURL)
  if err != nil {
    return nil
  }

  cfg.NextURL = response.Next
  cfg.PrevURL = response.Previous

  for _, result := range response.Results {
    fmt.Println(result.Name)
  }

  return nil
}

func commandMapB(cfg *Config, args... string) error {
  if cfg.PrevURL == nil {
    return nil
  }

  response, err := cfg.pokeApi.ListLocations(cfg.PrevURL)
  if err != nil {
    return nil
  }

  cfg.NextURL = response.Next
  cfg.PrevURL = response.Previous
  fmt.Println(cfg.NextURL, cfg.PrevURL)

  for _, result := range response.Results {
    fmt.Println(result.Name)
  }

  return nil
}


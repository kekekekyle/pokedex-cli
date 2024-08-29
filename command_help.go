package main

import (
  "fmt"
)

func commandHelp(cfg *Config, args... string) error {
  commands := getCommands()
  fmt.Println("")
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:")
  fmt.Println("")
  for _, command := range commands {
    fmt.Printf("%v: %v\n", command.name, command.description)
  }
  fmt.Println("")
  return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]CliCommand{
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
}

func main(){
  for {
    fmt.Printf("Pokedex > ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    command := scanner.Text()
    if c, ok := commands[command]; ok {
      c.callback()
    } else {
      fmt.Println("Command not found")
    }
  }
}

func commandExit() error {
	fmt.Println("Exiting program ....")
	os.Exit(1)
  return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

  fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
  return nil
}

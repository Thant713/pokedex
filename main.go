package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cfg := &config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				
			}
		},
	}
	repl(cfg)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type config struct {
	commands map[string]cliCommand
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "read error:", err)
				os.Exit(1)
			}
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		if cmd, ok := cfg.commands[words[0]]; ok {
			if err := cmd.callback(cfg); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for name, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

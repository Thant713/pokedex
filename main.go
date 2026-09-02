package main

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
				name:        "map",
				description: "Displays the next 20 location areas",
				callback:    commandMap,
			},
		},
	}
	repl(cfg)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
}

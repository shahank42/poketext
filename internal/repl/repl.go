package repl

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	Callback    func(args ...string) error
}

var CommandMap map[string]command

func Init() {
	CommandMap = map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback: func(args ...string) error {
				err := cmd_Help()
				return err
			},
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex application",
			Callback: func(args ...string) error {
				err := cmd_Exit()
				return err
			},
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 areas",
			Callback: func(args ...string) error {
				err := cmd_Map()
				return err
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas",
			Callback: func(args ...string) error {
				err := cmd_MapBack()
				return err
			},
		},
		"explore": {
			name:        "explore",
			description: "Shows the available Pokemon in a given area",
			Callback: func(args ...string) error {
				err := cmd_Explore(args[1])
				return err
			},
		},
		"catch": {
			name:        "catch",
			description: "Throws pokeball to a pokemon... can you catch it?",
			Callback: func(args ...string) error {
				err := cmd_Catch(args[1])
				return err
			},
		},
		"inspect": {
			name:        "inspect",
			description: "View details about a pokemon from your Pokedex",
			Callback: func(args ...string) error {
				err := cmd_Inspect(args[1])
				return err
			},
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your Pokedex",
			Callback: func(args ...string) error {
				err := cmd_Pokedex()
				return err
			},
		},
	}
}

func GetTextInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return scanner.Text()
}

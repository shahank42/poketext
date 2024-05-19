package main

import (
	"fmt"
	"strings"

	"github.com/shahank42/poketext/internal/app"
	"github.com/shahank42/poketext/internal/repl"
)

const PROMPT = "pokedex-cli >"

func main() {
	app.Init()
	repl.Init()

	for {
		fmt.Printf("%s ", PROMPT)
		userCmdInput := repl.GetTextInput()

		inputArgs := strings.Fields(userCmdInput)

		command, ok := repl.CommandMap[inputArgs[0]]
		if !ok {
			fmt.Println("Unknown command... need help? Type `help`")
			continue
		}

		err := command.Callback(inputArgs...)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

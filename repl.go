package main

import (
	"apso_todo/internal/database"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(db *database.DB) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("TO DO >")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(db, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*database.DB, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    commandExit,
		},
		"create": {
			name:        "create",
			description: "prompts user for name, and description then creates a todo item",
			callback:    commandCreate,
		},
		"complete": {
			name:        "complete",
			description: "enter the id of todo item to be completed as an argument to the command, ie: complete [id]",
			callback:    commandComplete,
		},
		"list": {
			name:        "list",
			description: "lists open todo items",
			callback:    commandList,
		},
	}
}

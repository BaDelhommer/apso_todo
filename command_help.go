package main

import (
	"apso_todo/internal/database"
	"fmt"
)

func commandHelp(db *database.DB, args ...string) error {
	cmds := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the APSO To Do app!")
	fmt.Println("Usage:")
	if len(args) == 1 {
		if cmd, ok := cmds[args[0]]; !ok {
			fmt.Println("unrecognized command", args[0])
		} else if ok {
			fmt.Println("Command: ", cmd.name)
			fmt.Println("Description: ", cmd.description)
		}
	} else {
		fmt.Println()
		for _, cmd := range cmds {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		fmt.Println()
	}
	return nil
}

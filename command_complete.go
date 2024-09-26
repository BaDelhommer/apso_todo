package main

import (
	"apso_todo/internal/database"
	"fmt"
	"strconv"
)

func commandComplete(db *database.DB, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Please provide an ID for a To Do item to complete")
	}

	if len(args) > 2 {
		fmt.Printf("Too many arguments, complete only takes 1 ID, got: %v\n", args[2])
	}

	for {

		id := args[0]
		intID, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		item, err := db.CompleteItem(intID)
		if err != nil {
			return err
		} else {
			fmt.Printf("Complted: %v, %s\n", item.ID, item.Name)
			break
		}
	}

	return nil

}

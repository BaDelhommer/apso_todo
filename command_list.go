package main

import (
	"apso_todo/internal/database"
	"fmt"
)

func commandList(db *database.DB, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("list takes no arguments, %v provided", args[1:])
	}

	items, err := db.ListOpenItems()
	if err != nil {
		return err
	}

	for _, item := range items {
		fmt.Println(item)
	}

	return nil
}

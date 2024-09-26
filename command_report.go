package main

import (
	"apso_todo/internal/database"
	"fmt"
)

func commandReport(db *database.DB, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("report takes no arguments, %v was given", args[1:])
	}

	openItems, err := db.ListOpenItems()
	if err != nil {
		return err
	}

	completeItems, err := db.ListCompleteItems()
	if err != nil {
		return err
	}

	fmt.Println(`
		=============
		ToDo Report
		=============
	`)

	fmt.Println("Open items:")
	for idx, item := range openItems {
		fmt.Printf("%v: %v\n", idx+1, item)
	}

	fmt.Println("Completed Items:")
	for idx, item := range completeItems {
		fmt.Printf("%v: %v\n", idx+1, item)
	}

	return nil
}

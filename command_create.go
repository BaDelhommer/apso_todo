package main

import (
	"apso_todo/internal/database"
	"bufio"
	"fmt"
	"os"
)

func commandCreate(db *database.DB, args ...string) error {
	if len(args) > 1 {
		fmt.Println("Create takes no arguments, got: ", args[1:])
	}
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter item name: ")
		reader.Scan()

		name := reader.Text()
		if len(name) == 0 {
			continue
		}

		fmt.Print("Enter item description: ")
		reader.Scan()

		desc := reader.Text()
		if len(desc) == 0 {
			continue
		}

		if len(name) == 0 && len(desc) == 0 {
			break
		}

		item, err := db.CreateItem(name, desc)
		if err != nil {
			return err
		}

		fmt.Println(item)
		break
	}
	return nil
}

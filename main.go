package main

import "apso_todo/internal/database"

func main() {
	db, err := database.NewDB("database.json")
	if err != nil {
		panic(err)
	}
	startRepl(db)
}

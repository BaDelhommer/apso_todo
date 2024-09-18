package main

import (
	"apso_todo/internal/database"
	"os"
)

func commandExit(db *database.DB, args ...string) error {
	os.Exit(0)
	return nil
}

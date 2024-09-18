package database

import "time"

type ToDoItem struct {
	ID           int
	Name         string
	Description  string
	DateAdded    time.Time
	DateComplete time.Time
}

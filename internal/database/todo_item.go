package database

import "time"

type ToDoItem struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	DateAdded    time.Time `json:"dateadded"`
	DateComplete time.Time `json:"datecompleted"`
}

func (db *DB) CreateItem(name, desc string) (ToDoItem, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return ToDoItem{}, err
	}

	id := len(dbStructure.Items) + 1
	item := ToDoItem{
		ID:          id,
		Name:        name,
		Description: desc,
		DateAdded:   time.Now().Local().UTC(),
	}

	dbStructure.Items[id] = item

	err = db.writeDB(dbStructure)
	if err != nil {
		return ToDoItem{}, nil
	}
	return item, nil
}

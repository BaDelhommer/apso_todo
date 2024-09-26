# apso_todo

#### Obligatory todo app

Built with Go version 1.22.2

Command line REPL to do app for no reason. I think there is a legal requirement that a person learning to code makes a to do app, so here is mine. I wanted something simple to track tasks as work, so I figured I would kill two birds with one stone.

Download the code and use ```go build``` to build the executable

Available commands:
*help: displays info about a command
*exit: exits the program
*create: prompts the user for a name and description, then creates a new item with that info
*complete: takes an id number of a todo item, then marks that item as complete
*list: lists open todo items
*report: prints a report to stdout showing all completed and open todo items
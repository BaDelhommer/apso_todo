# apso_todo

#### Obligatory todo app

Built with Go version 1.22.2

Command line REPL to do app for no reason. I think there is a legal requirement that a person learning to code makes a to do app, so here is mine. I wanted something simple to track tasks as work, so I figured I would kill two birds with one stone.

Download the code and use ```go build``` to build the executable

Running the program starts a REPL that takes commands.

Available commands:
* __help__: displays info about a commands, can take a command name as argument to display help for that command only
* __exit__: exits the program
* __create__: prompts the user for a name and description, then creates a new item with that info
* __complete__: takes an id number of a todo item, then marks that item as complete
* __list__: lists open todo items
* __report__: prints a report to stdout showing all completed and open todo items
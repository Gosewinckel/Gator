package main

import (
	"database/sql"
	"github.com/Gosewinckel/Gator/internal/database"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	s := initState()
	commands := initCommands(&s)

	db, err := sql.Open("postgres", s.conf.Db_url)
	if err != nil {
		fmt.Printf("could not open database: %v", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	s.db = dbQueries

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments.")
		os.Exit(1)
	} 
	cmdArgs := args[1: len(args)]
	command := command{Name: cmdArgs[0], Args: cmdArgs[1:len(cmdArgs)]}
	err = commands.run(&s, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

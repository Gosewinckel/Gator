package main

import (
	"fmt"
	"os"
)

func main() {
	s := initState()
	commands := initCommands(&s)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments.")
		os.Exit(1)
	}
	cmdArgs := args[1: len(args)]
	command := command{Name: cmdArgs[0], Args: cmdArgs[1:len(cmdArgs)]}
	err := commands.run(&s, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

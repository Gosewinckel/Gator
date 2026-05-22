package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		err := fmt.Errorf("Wrong number of commands. Shold only have username")
		return err
	}
	
	err := s.conf.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User: %v has been set.\n", cmd.Args[0])
	return nil
}

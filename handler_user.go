package main

import (
	"fmt"
	"os"
	"context"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		err := fmt.Errorf("Wrong number of commands. Should only have username")
		return err
	}

	ctx := context.Background()
	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		fmt.Println("user could not be found")
		os.Exit(1)
	}
	
	err = s.conf.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User: %v has been set.\n", cmd.Args[0])
	return nil
}

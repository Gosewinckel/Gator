package main

import (
	"fmt"
	"context"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		err := fmt.Errorf("only need reset keyword")
		return err
	}

	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	
	// Print all users
	for name := range(users) {
		if users[name] == s.conf.Current_user_name {
			fmt.Printf("* %v (current)\n", users[name])
		} else {
			fmt.Printf("* %v\n", users[name])
		}
	}
	return nil
}

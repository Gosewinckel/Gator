package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/google/uuid"
	"github.com/Gosewinckel/Gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		err := fmt.Errorf("wron number of args, must have a name")
		return err
	}

	ctx := context.Background()
	id := uuid.New()

	current_time := time.Now()
	name := cmd.Args[0]

	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		os.Exit(1)
	}

	params := database.CreateUserParams{id, current_time, current_time, name}
	user, err := s.db.CreateUser(ctx, params)
	
	s.conf.SetUser(user.Name)
	fmt.Printf("User was created, user data: %v\n", user)
	return nil
}

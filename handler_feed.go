package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/Gosewinckel/Gator/internal/database"
)

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		err := fmt.Errorf("Command must have 2 arguments")
		return err
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.conf.Current_user_name)
	if err != nil {
		return err
	}

	id := uuid.New()
	currTime := time.Now()
	feedVals := database.CreateFeedParams{id, currTime, currTime, cmd.Args[0], cmd.Args[1], user.ID}
	feed, err := s.db.CreateFeed(ctx, feedVals)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil

}

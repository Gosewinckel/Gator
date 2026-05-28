package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/Gosewinckel/Gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		err := fmt.Errorf("wrong number of args for this command")
		return err
	}

	ctx := context.Background()
	feed, err := s.db.GetFeed(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	user, err := s.db.GetUser(ctx, s.conf.Current_user_name)
	if err != nil {
		return err
	} 

	current := time.Now()
	id := uuid.New()

	params := database.CreateFeedFollowsParams{}
	params.ID = id
	params.CreatedAt = current
	params.UpdatedAt = current
	params.FeedID = feed.ID
	params.UserID = user.ID

	new_feed_follow, err := s.db.CreateFeedFollows(ctx, params)
	if err != nil {
		return err
	}

	for i := range(new_feed_follow) {
		fmt.Println("Feed: %v, User: %v\n", new_feed_follow[i].FeedName, new_feed_follow[i].UserName)
	}

	return nil
}

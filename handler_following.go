package main

import (
	"fmt"
	"context"

	"github.com/Gosewinckel/Gator/internal/database"
)

func handlerFollowing(s *state, cmd command, use database.User) error {
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.conf.Current_user_name)
	if err != nil {
		return err
	}

	userId := user.ID
	userFeed, err := s.db.GetFeedFollowsForUser(ctx, userId)
	if err != nil {
		return err
	}

	for i := range(userFeed) {
		fmt.Printf("Feed Name: %v\n", userFeed[i].Name_2)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, use database.User) error {
	if len(cmd.Args) != 1 {
		err := fmt.Errorf("wrong number of args")
		return err
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.conf.Current_user_name)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowsParams{user.ID, feed.ID}
	err = s.db.DeleteFeedFollows(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"context"
)

func handlerFollowing(s *state, cmd command) error {
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

package main

import (
	"fmt"
	"context"
)

func handlerAggregator(s *state, cmd command) error {
	ctx := context.Background()
	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	
	fmt.Println(feed)
	return nil
}

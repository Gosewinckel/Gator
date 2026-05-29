package main

import (
	"fmt"
	"context"
	"time"

	"github.com/Gosewinckel/Gator/internal/database"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		err := fmt.Errorf("Wrong numer of arguments")
		return err
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feed every", interval)
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}
	
	current := time.Now()
	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{nextFeed.ID, current})	
	if err != nil {
		return err
	}

	feed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println(feed.Channel.Title)
	for i := range(feed.Channel.Item) {
		fmt.Println(feed.Channel.Item[i].Title)
	}
	return nil
}

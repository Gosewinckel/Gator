package main
import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		err := fmt.Errorf("Too many arguments")
		return err
	}

	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}
	for i := range(len(feeds)) {
		fmt.Printf("Name: %v, URL: %v, User name: %v\n", feeds[i].Name, feeds[i].Url, feeds[i].Name_2)
	}
	return nil
}

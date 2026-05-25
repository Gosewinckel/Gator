package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		err := fmt.Errorf("should not be additional arguments")
		os.Exit(1)
		return err
	}

	ctx := context.Background()
	err := s.db.ResetUsers(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}

	return nil
}

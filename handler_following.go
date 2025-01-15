package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, _ command) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Feed)
	}

	return nil
}

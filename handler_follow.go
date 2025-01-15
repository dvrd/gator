package main

import (
	"context"
	"fmt"
	"github.com/dvrd/gator/internal/database"
	"github.com/google/uuid"
	"strings"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: gator follow <url>")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(context.Background(), strings.TrimSpace(cmd.Args[0]))
	if err != nil {
		return err
	}

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), newFollow)
	if err != nil {
		return err
	}

	fmt.Println(follow.FeedName)
	fmt.Println(follow.UserName)

	return nil
}

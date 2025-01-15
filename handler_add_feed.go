package main

import (
	"context"
	"fmt"
	"github.com/dvrd/gator/internal/database"
	"github.com/google/uuid"
	"reflect"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Usage: gator addfeed <name> <url>")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}

	err = handlerFollow(s, command{
		Name: "follow",
		Args: []string{feed.Url},
	})
	if err != nil {
		return err
	}

	v := reflect.ValueOf(feed)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}

	return nil
}

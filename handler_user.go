package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dvrd/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func parseError(e error) error {
	pgErr, ok := e.(*pq.Error)

	if ok {
		switch pgErr.Code.Name() {
		case "unique_violation":
			return fmt.Errorf("User already exists")
		case "undefined_table":
			return fmt.Errorf("Users table does not exist")
		}
	}

	return e
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: gator register <name>")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return parseError(err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User successfuly created")
	fmt.Println(user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: gator login <name>")
	}

	name := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("User is not registered yet")
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

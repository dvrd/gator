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

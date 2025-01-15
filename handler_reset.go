package main

import (
	"context"
)

func handlerReset(s *state, cmd command) error {
	return s.db.DeleteUsers(context.Background())
}

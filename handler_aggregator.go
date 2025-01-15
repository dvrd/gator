package main

import (
	"context"
	"fmt"
)

const tmpUrl = "https://www.wagslane.dev/index.xml"

func handlerAggregator(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), tmpUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Channel: \033[1m%s\033[0m\n", feed.Channel.Title)
	fmt.Printf("Description: \033[1m%s\033[0m\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("\n\033[1m%s\033[0m\n", item.Title)
		fmt.Printf("\033[32m%s\033[0m\n\n", item.Link)
		fmt.Println(item.Description)
	}

	return nil
}

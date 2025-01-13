package main

import (
	"fmt"
	"github.com/dvrd/gator/internal/config"
	"os"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println("Could not read config file:", err)
		os.Exit(1)
	}

	fmt.Println(config.DBUrl)
}

package main

import (
	"fmt"
	"github.com/dvrd/gator/internal/config"
	"os"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = config.SetUser("dvrd")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*config)
}

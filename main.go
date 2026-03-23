package main

import (
	"fmt"
	"log"

	"gator/internal/config"
)

func main() {
	// 1. Read the config file from disk
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config before:", cfg)

	// 2. Set the current user and write the updated config to disk
	if err := cfg.SetUser("ero"); err != nil {
		log.Fatal(err)
	}

	// 3. Read the config file again to confirm the change was saved
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config after:", cfg)
}

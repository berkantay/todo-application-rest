package main

import (
	"context"
	"log"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/repository"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewConfig(context.Background(), "development", ".config")
	if err != nil {
		log.Fatal(err)
	}

	_, err = repository.NewDatabase(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
}

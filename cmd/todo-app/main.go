package main

import (
	"context"
	"log"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/database"
	"github.com/berkantay/todo-app-example/repository"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewConfig(context.Background(), "development", ".config")
	if err != nil {
		log.Fatal(err)
	}

	database, err := database.NewDatabase(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(*database)

	// userService := user.NewService(context.Background(), userRepository)

}

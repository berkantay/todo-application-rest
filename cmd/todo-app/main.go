package main

import (
	"context"
	"log"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/database"
	"github.com/berkantay/todo-app-example/repository"
	"github.com/berkantay/todo-app-example/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// authService := auth.NewService(context.Background(), userRepository, config)
	userService := user.NewService(context.Background(), userRepository)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	api := app.Group("/api")
	user.RegisterRouters(api, *userService)
	log.Fatal(app.Listen(":8081"))
}

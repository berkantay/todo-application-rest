package main

import (
	"context"
	"log"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/db"
	"github.com/berkantay/todo-app-example/internal/router"
	"github.com/berkantay/todo-app-example/internal/todo"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewConfig(context.Background(), "development", ".config")
	if err != nil {
		log.Fatal(err)

	}

	dbConn, err := db.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	todoPersistanceTransaction := port.NewPostgresTodoPort(dbConn.Instance())

	todoRepository := todo.NewRepository(todoPersistanceTransaction)
	todoService := todo.NewService(todoRepository)
	todoHandler := todo.NewHandler(todoService)

	router.SetupRouter(todoHandler)
	router.Run("0.0.0.0:8080")
}

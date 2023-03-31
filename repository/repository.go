package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/repository/ent"
)

type Database struct {
	Client *ent.Client
}

func NewDatabase(ctx context.Context, config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Postgresql.Url,
		config.Postgresql.Port,
		config.Postgresql.Username,
		config.Postgresql.Name,
		config.Postgresql.Password,
	)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &Database{
		Client: client,
	}, nil
}

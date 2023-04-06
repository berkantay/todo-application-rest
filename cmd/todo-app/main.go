package main

import (
	"context"
	"fmt"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/db"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewConfig(context.Background(), "development", ".config")
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.NewDatabase(config)
	if err != nil {
		fmt.Println(err)
	}

}

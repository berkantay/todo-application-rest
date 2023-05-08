package db

import (
	"database/sql"
	"fmt"

	"github.com/berkantay/todo-app-example/config"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/todos?sslmode=disable",
		config.Postgresql.Username,
		config.Postgresql.Password,
		config.Postgresql.Url,
		config.Postgresql.Port,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() error {
	err := d.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Instance() *sql.DB {
	return d.db

}

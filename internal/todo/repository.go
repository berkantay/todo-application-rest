package todo

import (
	"context"
	"database/sql"

	"github.com/berkantay/todo-app-example/internal/entity"
)

type DatabaseTransaction interface {
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Repository interface {
	CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
}

type repository struct {
	db DatabaseTransaction
}

func NewRepository(db DatabaseTransaction) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	var lastInsertId int
	query := "INSERT INTO todos(description, deadline, priority) VALUES ($1, $2 ,$3) returning id"
	err := r.db.QueryRowContext(ctx, query, todo.Description, todo.Deadline, todo.Priority).Scan(&lastInsertId)
	if err != nil {
		return &entity.Todo{}, err
	}
	todo.Id = int64(lastInsertId)
	return todo, nil
}

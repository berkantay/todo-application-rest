package todo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/berkantay/todo-app-example/internal/entity"
)

type DatabaseTransactionPort interface {
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type PostgresTodoAdapter struct {
	tx DatabaseTransactionPort
}

func NewPostgresTodoAdapter(dbConn DatabaseTransactionPort) *PostgresTodoAdapter {
	return &PostgresTodoAdapter{
		tx: dbConn,
	}
}

func (ptp *PostgresTodoAdapter) CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	var lastInsertId int
	query := "INSERT INTO todos(description, deadline, priority) VALUES ($1, $2 ,$3) returning id"
	err := ptp.tx.QueryRowContext(ctx, query, todo.Description, todo.Deadline, todo.Priority).Scan(&lastInsertId)
	if err != nil {
		return &entity.Todo{}, err
	}
	todo.Id = int64(lastInsertId)
	return todo, nil
}
func (ptp *PostgresTodoAdapter) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	allTodo := make([]*entity.Todo, 0)
	query := "SELECT * FROM todos"
	rows, err := ptp.tx.QueryContext(ctx, query)
	if err != nil {
		return allTodo, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo entity.Todo
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Deadline, &todo.Priority); err != nil {
			return nil, err
		}
		allTodo = append(allTodo, &todo)
	}
	return allTodo, nil
}

func (ptp *PostgresTodoAdapter) DeleteTodo(ctx context.Context, id int) (*int, error) {
	query := "DELETE FROM todos WHERE id = $1;"
	res, err := ptp.tx.ExecContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affectedRows == 0 {
		return nil, errors.New("no entry found")
	}
	return &id, nil
}

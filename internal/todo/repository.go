package todo

import (
	"context"

	"github.com/berkantay/todo-app-example/internal/entity"
)

type PostgreTodoRepository interface {
	CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
	DeleteTodo(ctx context.Context, id int) (*int, error)
	GetAllTodos(ctx context.Context) ([]*entity.Todo, error)
}

type repository struct {
	db PostgreTodoRepository
}

func NewRepository(db PostgreTodoRepository) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	return r.db.CreateTodo(ctx, todo)
}
func (r *repository) Delete(ctx context.Context, id int) (*int, error) {
	return r.db.DeleteTodo(ctx, id)
}
func (r *repository) GetAll(ctx context.Context) ([]*entity.Todo, error) {
	return r.db.GetAllTodos(ctx)
}

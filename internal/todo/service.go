package todo

import (
	"context"
	"errors"
	"time"

	"github.com/berkantay/todo-app-example/internal/entity"
)

const (
	maxPriority = 3
)

type service struct {
	todoRepository Repository
	timeout        time.Duration
}

type Service interface {
	CreateTodo(context.Context, *entity.Todo) (*entity.Todo, error)
	GetAllTodo(context.Context) ([]*entity.Todo, error)
	DeleteTodo(context.Context, int) (*entity.Todo, error)
}

func NewService(repository Repository) Service {
	return &service{
		todoRepository: repository,
		timeout:        time.Duration(2 * time.Second),
	}
}

func (s *service) CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if !checkPriorityScale(todo.Priority) {
		return &entity.Todo{}, errors.New("invalid priority range")
	}

	t := &entity.Todo{
		Description: todo.Description,
		Deadline:    todo.Deadline,
		Priority:    todo.Priority,
	}

	createdTodo, err := s.todoRepository.Create(ctx, t)
	if err != nil {
		return &entity.Todo{}, err
	}
	t.Id = createdTodo.Id
	return t, nil
}

func (s *service) DeleteTodo(ctx context.Context, id int) (*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	deleteTodoId, err := s.todoRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Todo{
		Id: int64(*deleteTodoId),
	}, nil
}

func (s *service) GetAllTodo(ctx context.Context) ([]*entity.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	todos, err := s.todoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func checkPriorityScale(priority int64) bool {
	if priority > 0 || priority <= maxPriority {
		return true
	}
	return false
}

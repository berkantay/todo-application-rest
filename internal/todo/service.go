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
	repository Repository
	timeout    time.Duration
}

type Service interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*entity.Todo, error)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		timeout:    time.Duration(2 * time.Second),
	}
}

type CreateTodoRequest struct {
	Description string `json:"description,omitempty"`
	Deadline    string `json:"deadline,omitempty"`
	Priority    int64  `json:"priority,omitempty"` //0-3 scale
}

func (s *service) CreateTodo(ctx context.Context, todo *CreateTodoRequest) (*entity.Todo, error) {
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

	createdTodo, err := s.repository.CreateTodo(ctx, t)
	if err != nil {
		return &entity.Todo{}, err
	}
	t.Id = createdTodo.Id
	return t, nil
}

func checkPriorityScale(priority int64) bool {
	if priority > 0 || priority <= maxPriority {
		return true
	}
	return false
}

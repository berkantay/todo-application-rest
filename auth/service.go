package auth

import (
	"context"
	"errors"

	"github.com/berkantay/todo-app-example/config"
	"github.com/berkantay/todo-app-example/model"
)

type Repository interface {
	ReadUser(ctx context.Context, username, password string) ([]*model.User, error)
}

type Service struct {
	authRepository Repository
}

func NewService(ctx context.Context, authRepository Repository, config *config.Config) *Service {
	return &Service{
		authRepository: authRepository,
	}
}

func (s *Service) Authenticate(ctx context.Context, username, password string) error {
	users, err := s.authRepository.ReadUser(ctx, username, password)
	if err != nil {
		return err
	}

	if users == nil {
		return errors.New("user not found")
	}
	return nil
}

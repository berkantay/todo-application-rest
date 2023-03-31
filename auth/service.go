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
	userRepository Repository
	config         *config.Config
}

func NewService(ctx context.Context, userRepository Repository, config *config.Config) *Service {
	return &Service{
		userRepository: userRepository,
		config:         config,
	}
}

func (s *Service) Authenticate(ctx context.Context, username, password string) error {
	users, err := s.userRepository.ReadUser(ctx, username, password)
	if err != nil {
		return err
	}

	if users == nil {
		return errors.New("user not found")
	}
	return nil
}

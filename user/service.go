package user

import (
	"context"
	"errors"

	"github.com/berkantay/todo-app-example/model"
)

const (
	minPasswordLength = 8
)

type Repository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Read(ctx context.Context, id string) ([]*model.User, error)
	UpdateUsername(ctx context.Context, id string, username string) (*model.User, error)
	UpdatePassword(ctx context.Context, id string, password string) (*model.User, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	userRepository Repository
}

func NewService(ctx context.Context, repository Repository) *Service {
	return &Service{
		userRepository: repository,
	}
}

func (s *Service) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if !isValidUsername(*user.Username) {
		return nil, errors.New("username cannot be empty")
	}

	err := isValidPassword(*user.Password)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *Service) Read(ctx context.Context, id string) ([]*model.User, error) {
	users, err := s.userRepository.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) UpdateUsername(ctx context.Context, id string, username string) (*model.User, error) {
	if !isValidUsername(username) {
		return nil, errors.New("username cannot be empty")
	}
	user, err := s.userRepository.UpdateUsername(ctx, id, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UpdatePassword(ctx context.Context, id string, password string) (*model.User, error) {
	err := isValidPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepository.UpdatePassword(ctx, id, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func isValidUsername(username string) bool {
	return username != ""
}

func isValidPassword(password string) error {
	if len(password) < minPasswordLength {
		return errors.New("password should be longer than 8 characters")
	}
	return nil
}

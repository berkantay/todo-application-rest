package user

import (
	"context"

	"github.com/berkantay/todo-app-example/model"
)

type Repository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Read(ctx context.Context, id string) ([]*model.User, error)
	UpdateUsername(ctx context.Context, id string, username string) (*model.User, error)
	UpdatePassword(ctx context.Context, id string, password string) (*model.User, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repository Repository
}

func NewService(ctx context.Context, repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *Service) Read(ctx context.Context, id string) ([]*model.User, error) {
	users, err := s.repository.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) UpdateUsername(ctx context.Context, id string, username string) (*model.User, error) {
	user, err := s.repository.UpdateUsername(ctx, id, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UpdatePassword(ctx context.Context, id string, password string) (*model.User, error) {
	user, err := s.repository.UpdatePassword(ctx, id, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

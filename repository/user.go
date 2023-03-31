package repository

import (
	"context"

	repository "github.com/berkantay/todo-app-example/database"
	"github.com/berkantay/todo-app-example/database/ent"
	"github.com/berkantay/todo-app-example/database/ent/user"
	"github.com/berkantay/todo-app-example/model"
	"github.com/google/uuid"
)

type UserRepository struct {
	instance repository.Database
}

func NewUserRepository(instance repository.Database) *UserRepository {
	return &UserRepository{
		instance: instance,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := ur.instance.Client.User.Create().SetUsername(*user.Username).SetPassword(*user.Password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return toUserModel(createdUser), nil
}

func (ur *UserRepository) Read(ctx context.Context, id string) ([]*model.User, error) {
	queried, err := ur.instance.Client.User.Query().Where(user.ID(uuid.MustParse(id))).All(ctx)
	if err != nil {
		return nil, err
	}
	return userEntityToModel(queried), nil
}

func (ur *UserRepository) UpdateUsername(ctx context.Context, id string, username string) (*model.User, error) {
	updatedUser, err := ur.instance.Client.User.UpdateOneID(uuid.MustParse(id)).SetUsername(username).Save(ctx)
	if err != nil {
		return nil, err
	}
	return toUserModel(updatedUser), nil
}

func (ur *UserRepository) UpdatePassword(ctx context.Context, id string, password string) (*model.User, error) {
	updatedUser, err := ur.instance.Client.User.UpdateOneID(uuid.MustParse(id)).SetPassword(password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return toUserModel(updatedUser), nil
}

func (ur *UserRepository) ReadUser(ctx context.Context, username, password string) ([]*model.User, error) {
	queried, err := ur.instance.Client.User.Query().Where(user.Username(username)).Where(user.Password(password)).All(ctx)
	if err != nil {
		return nil, err
	}
	return userEntityToModel(queried), nil
}

func (ur *UserRepository) Delete(ctx context.Context, id string) error {
	err := ur.instance.Client.User.DeleteOneID(uuid.MustParse(id)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func toUserModel(queried *ent.User) *model.User {
	return &model.User{
		UUID:     queried.ID,
		Username: &queried.Username,
		Password: &queried.Password,
	}
}

func userEntityToModel(queried []*ent.User) []*model.User {
	users := make([]*model.User, 0)
	for _, q := range queried {
		users = append(users, &model.User{
			UUID:     q.ID,
			Username: &q.Username,
			Password: &q.Password,
		})
	}
	return users
}

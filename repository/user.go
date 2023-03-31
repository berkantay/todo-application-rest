package repository

import (
	"context"

	"github.com/berkantay/todo-app-example/model"
	"github.com/berkantay/todo-app-example/repository/ent"
	"github.com/berkantay/todo-app-example/repository/ent/user"
	"github.com/google/uuid"
)

func (d *Database) Create(ctx context.Context, user *model.User) (*model.User, error) {
	createdUser, err := d.Client.User.Create().SetUsername(*user.Username).SetPassword(*user.Password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return toUserModel(createdUser), nil
}

func (d *Database) Read(ctx context.Context, id string) ([]*model.User, error) {
	queried, err := d.Client.User.Query().Where(user.ID(uuid.MustParse(id))).All(ctx)
	if err != nil {
		return nil, err
	}
	return userEntityToModel(queried), nil
}

func (d *Database) UpdateUsername(ctx context.Context, id string, username string) (*model.User, error) {
	updatedUser, err := d.Client.User.UpdateOneID(uuid.MustParse(id)).SetUsername(username).Save(ctx)
	if err != nil {
		return nil, err
	}
	return toUserModel(updatedUser), nil
}

func (d *Database) UpdatePassword(ctx context.Context, id string, password string) (*model.User, error) {
	updatedUser, err := d.Client.User.UpdateOneID(uuid.MustParse(id)).SetPassword(password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return toUserModel(updatedUser), nil
}

func (d *Database) Delete(ctx context.Context, id string) error {

	err := d.Client.User.DeleteOneID(uuid.MustParse(id)).Exec(ctx)
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

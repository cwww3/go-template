package user

import (
	"context"
	"fmt"
	"github.com/cwww3/go-template/internal/entity"
	"github.com/cwww3/go-template/internal/repository"
)

const userfield = "name,email,phone"

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	db repository.DBTX
}

func NewUserRepository(db repository.DBTX) repository.UserRepository {
	return userRepository{db: db}
}

func (ur userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	result, err := ur.db.ExecContext(ctx, fmt.Sprintf("insert into user(%s) values(?,?,?)", userfield), user.Name, user.Email, user.Phone)
	if err != nil {
		return nil, err
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur userRepository) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return nil, nil
}

func (ur userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (ur userRepository) Delete(ctx context.Context, id int64) error {
	_, err := ur.db.ExecContext(ctx, "delete from user where id = ?", id)
	return err
}

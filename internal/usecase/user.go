package usecase

import (
	"context"
	"github.com/cwww3/go-template/internal/entity"
	"github.com/cwww3/go-template/internal/repository"
)

type UserUseCase interface {
	AddUser(context.Context, *entity.User) (*entity.User, error)
	AddUserAndDelete(context.Context, *entity.User) error
	GetUser(context.Context, int64) (*entity.User, error)
	ModifyUser(context.Context, *entity.User) (*entity.User, error)
}

type userUseCase struct {
	repository repository.Repository
}

func (u *userUseCase) AddUserAndDelete(ctx context.Context, user *entity.User) error {
	return u.repository.Atomic(ctx, func(r repository.Repository) error {
		u, err := r.GetUserRepository().CreateUser(ctx, user)
		if err != nil {
			return err
		}
		return r.GetUserRepository().Delete(ctx, u.ID)
	})
}

func NewUserUseCase(r repository.Repository) UserUseCase {
	return &userUseCase{
		repository: r,
	}
}

func (u *userUseCase) AddUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.repository.GetUserRepository().CreateUser(ctx, user)
}

func (u *userUseCase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return u.repository.GetUserRepository().GetUser(ctx, id)
}

func (u *userUseCase) ModifyUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.repository.GetUserRepository().UpdateUser(ctx, user)
}

package user

import (
	"context"
	"github.com/cwww3/go-template/internal/entity"
	"github.com/cwww3/go-template/internal/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return userRepository{db: db}
}

func (ur userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return user, ur.db.Model(new(entity.User)).Create(user).Error
}

func (ur userRepository) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	return &user, ur.db.Model(new(entity.User)).First(&user, id).Error
}

func (ur userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return user, ur.db.Model(new(entity.User)).Updates(user).Error
}

func (ur userRepository) Delete(ctx context.Context, id int64) error {
	return ur.db.Where("id = ?", id).Delete(new(entity.User)).Error
}

package orm

import (
	"context"
	"github.com/cwww3/go-template/internal/repository"
	"github.com/cwww3/go-template/internal/repository/orm/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ormRepository struct {
	db *gorm.DB
}

func NewOrmRepository(dsn string) repository.Repository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic("failed to connect database")
	}
	return ormRepository{
		db: db,
	}
}

func (or ormRepository) GetUserRepository() repository.UserRepository {
	return user.NewUserRepository(or.db)
}

func (or ormRepository) Atomic(ctx context.Context, fn func(r repository.Repository) error) error {
	return or.db.Transaction(func(tx *gorm.DB) error {
		return fn(&ormRepository{db: tx})
	})
}

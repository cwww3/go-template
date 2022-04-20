package orm

import (
	"context"
	"fmt"
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
	tx := or.db.Begin()
	var err error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			switch p.(type) {
			case error:
				err = fmt.Errorf("panic err: %v", p)
			default:
				panic(err)
			}
		}
	}()
	err = fn(&ormRepository{db: tx})
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

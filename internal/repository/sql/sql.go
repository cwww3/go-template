package sql

import (
	"context"
	"database/sql"

	"fmt"
	"github.com/cwww3/go-template/internal/repository"
	"github.com/cwww3/go-template/internal/repository/sql/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sqlRepository struct {
	db         *sqlx.DB
	dbexecutor repository.DBTX
}

func (sr sqlRepository) GetUserRepository() repository.UserRepository {
	return user.NewUserRepository(sr.dbexecutor)
}

func NewSqlRepository(dsn string) repository.Repository {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return sqlRepository{db: db, dbexecutor: db}
}

func (sr sqlRepository) Atomic(ctx context.Context, fn func(r repository.Repository) error) error {
	tx, err := sr.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
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
	err = fn(&sqlRepository{dbexecutor: tx})
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

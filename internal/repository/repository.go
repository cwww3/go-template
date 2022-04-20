package repository

import (
	"context"
	"database/sql"
	"github.com/cwww3/go-template/internal/entity"
)

type Repository interface {
	Atomic(ctx context.Context, fn func(r Repository) error) error
	GetUserRepository() UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

// DBTX sql.db 和 sql.tx 的公共接口
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

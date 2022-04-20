package sql

import (
	"context"
	"errors"
	"github.com/cwww3/go-template/internal/entity"
	"github.com/cwww3/go-template/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newRepository() repository.Repository {
	dsn := "root:12345678@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	return NewSqlRepository(dsn)
}

var CustomErr = errors.New("custom err")

func TestAtomicSuccess(t *testing.T) {
	r := newRepository()
	ctx := context.Background()
	user := &entity.User{
		Name:  "cw",
		Email: "cw@qq.com",
		Phone: "188",
	}
	err := r.Atomic(ctx, func(r repository.Repository) error {
		t.Log("create user")
		user, err := r.GetUserRepository().CreateUser(ctx, user)
		if err != nil {
			return err
		}
		t.Logf("user %v\n", *user)

		// make err
		//return errors.New("custom err")

		t.Log("delete user")
		err = r.GetUserRepository().Delete(ctx, user.ID)
		return err
	})
	if !assert.Equal(t, nil, err) {
		t.FailNow()
	}
}

func TestAtomicFailed(t *testing.T) {
	r := newRepository()
	ctx := context.Background()
	user := &entity.User{
		Name:  "cw",
		Email: "cw@qq.com",
		Phone: "188",
	}
	err := r.Atomic(ctx, func(r repository.Repository) error {
		t.Log("create user")
		user, err := r.GetUserRepository().CreateUser(ctx, user)
		if err != nil {
			return err
		}
		t.Logf("user %v\n", *user)

		//make err
		return CustomErr

		t.Log("delete user")
		err = r.GetUserRepository().Delete(ctx, user.ID)
		return err
	})
	if !assert.Equal(t, CustomErr, err) {
		t.FailNow()
	}
}

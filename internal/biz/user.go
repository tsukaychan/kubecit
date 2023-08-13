/**
 * @author tsukiyo
 * @date 2023-08-13 3:00
 */

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
	Password string
	Age      int
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Register(context.Context, *User) (*User, error)
	List(ctx context.Context) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *UserUsecase) Register(ctx context.Context, user *User) (*User, error) {
	userResult, err := u.repo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (u *UserUsecase) List(ctx context.Context) ([]*User, error) {
	userResult, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

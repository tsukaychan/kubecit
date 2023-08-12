/**
 * @author tsukiyo
 * @date 2023-08-13 2:54
 */

package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) Register(ctx context.Context, user *biz.User) (*biz.User, error) {
	userEnt, err := u.data.db.User.Create().SetName(user.Username).SetAge(1).SetPassword(user.Password).Save(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return &biz.User{
		Username: userEnt.Name,
		Password: userEnt.Password,
		Age:      userEnt.Age,
	}, nil
}

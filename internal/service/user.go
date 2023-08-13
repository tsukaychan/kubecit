package service

import (
	"context"
	"fmt"

	v1 "kubecit/api/helloworld/v1"
	"kubecit/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServiceServer
	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(userCase *biz.UserUsecase) *UserService {
	return &UserService{uc: userCase}
}

// UserRegister register a user with username and password
func (s *UserService) UserRegister(ctx context.Context, in *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	fmt.Println(in.Username, in.Password)
	_, err := s.uc.Register(ctx, &biz.User{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{Result: "success"}, nil
}

func (s *UserService) UserList(ctx context.Context, in *v1.Empty) (*v1.UserListResponse, error) {
	userList, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	userRes := []*v1.User{}
	for _, user := range userList {
		userRes = append(userRes, &v1.User{
			Username: user.Username,
			Password: user.Password,
		})
	}
	return &v1.UserListResponse{
		Users: userRes,
	}, nil
}

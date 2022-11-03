package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// UserLogin user login,token in api gateway
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success

	err = user.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}

	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

func (*UserService) UserRegister(ctx context.Context, request *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user repository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Success

	err = user.UserCreate(request)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}

	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

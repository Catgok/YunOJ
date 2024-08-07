// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"YunOJ/services/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoByIdRequest  = user.GetUserInfoByIdRequest
	GetUserInfoByIdResponse = user.GetUserInfoByIdResponse
	LoginByUserKeyRequest   = user.LoginByUserKeyRequest
	LoginByUserKeyResponse  = user.LoginByUserKeyResponse
	RegisterRequest         = user.RegisterRequest
	RegisterResponse        = user.RegisterResponse
	User                    = user.User

	UserService interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		LoginByUserKey(ctx context.Context, in *LoginByUserKeyRequest, opts ...grpc.CallOption) (*LoginByUserKeyResponse, error)
		GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) LoginByUserKey(ctx context.Context, in *LoginByUserKeyRequest, opts ...grpc.CallOption) (*LoginByUserKeyResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.LoginByUserKey(ctx, in, opts...)
}

func (m *defaultUserService) GetUserInfoById(ctx context.Context, in *GetUserInfoByIdRequest, opts ...grpc.CallOption) (*GetUserInfoByIdResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfoById(ctx, in, opts...)
}

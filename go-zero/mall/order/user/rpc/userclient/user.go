// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	UserResponse = user.UserResponse
	IdRequest    = user.IdRequest

	User interface {
		GetUser(ctx context.Context, in *IdRequest) (*UserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in)
}

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"CrabGo/internal/model"
	"CrabGo/internal/model/entity"
	"context"
)

type (
	IUser interface {
		Auth(ctx context.Context, input model.UserAuthInput) (out *model.UserAuthOutput, err error)
		GetById(ctx context.Context, userId string) (user *entity.SysUser, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}

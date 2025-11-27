// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"CrabGo/internal/model"
	"context"
)

type (
	IMachine interface {
		QueryCpuInfo(ctx context.Context, input *model.QueryCpuInfoInput) (output *model.QueryCpuInfoOutput, err error)
	}
)

var (
	localMachine IMachine
)

func Machine() IMachine {
	if localMachine == nil {
		panic("implement not found for interface IMachine, forgot register?")
	}
	return localMachine
}

func RegisterMachine(i IMachine) {
	localMachine = i
}

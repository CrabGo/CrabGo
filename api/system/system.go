// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"CrabGo/api/system/v1"
)

type ISystemV1 interface {
	QueryCpus(ctx context.Context, req *v1.QueryCpusReq) (res *v1.QueryCpusRes, err error)
	QueryCpuRate(ctx context.Context, req *v1.QueryCpuRateReq) (res *v1.QueryCpuRateRes, err error)
	QueryMemory(ctx context.Context, req *v1.QueryMemoryReq) (res *v1.QueryMemoryRes, err error)
	QueryDisk(ctx context.Context, req *v1.QueryDiskReq) (res *v1.QueryDiskRes, err error)
	QueryHost(ctx context.Context, req *v1.QueryHostReq) (res *v1.QueryHostRes, err error)
	QueryNet(ctx context.Context, req *v1.QueryNetReq) (res *v1.QueryNetRes, err error)
	QueryProcess(ctx context.Context, req *v1.QueryProcessReq) (res *v1.QueryProcessRes, err error)
}

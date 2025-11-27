// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package build

import (
	"context"

	"CrabGo/api/build/v1"
)

type IBuildV1 interface {
	BuildInfo(ctx context.Context, req *v1.BuildInfoReq) (res *v1.BuildInfoRes, err error)
}

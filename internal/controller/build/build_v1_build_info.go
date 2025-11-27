package build

import (
	"context"

	"CrabGo/api/build/v1"

	"github.com/gogf/gf/v2/os/gbuild"
)

func (c *ControllerV1) BuildInfo(ctx context.Context, req *v1.BuildInfoReq) (res *v1.BuildInfoRes, err error) {

	buildInfo := gbuild.Info()
	res = &v1.BuildInfoRes{
		Version: buildInfo.Version,
		GF:      buildInfo.GoFrame,
		Go:      buildInfo.Golang,
		Git:     buildInfo.Git,
		Time:    buildInfo.Time,
	}

	return res, nil
}

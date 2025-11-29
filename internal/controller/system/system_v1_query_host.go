package system

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v4/host"

	"CrabGo/api/system/v1"
)

func (c *ControllerV1) QueryHost(ctx context.Context, req *v1.QueryHostReq) (res *v1.QueryHostRes, err error) {

	res = &v1.QueryHostRes{}

	hostInfo, err := host.Info()

	if err != nil {
		return nil, err
	}

	gconv.Struct(hostInfo, &res)

	return res, nil
}

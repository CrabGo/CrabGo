package machine

import (
	"context"

	"CrabGo/api/machine/v1"

	"github.com/shirou/gopsutil/cpu"
)

func (c *ControllerV1) QueryCpuRate(ctx context.Context, req *v1.QueryCpuRateReq) (res *v1.QueryCpuRateRes, err error) {

	res = &v1.QueryCpuRateRes{}

	precentValue, err := cpu.Percent(0, false)

	if err != nil {
		return nil, err
	}

	res.Use = precentValue[0]

	precentValue, err = cpu.Percent(0, true)

	if err != nil {
		return nil, err
	}

	res.CoreUser = precentValue

	return res, nil
}

package machine

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/cpu"

	"CrabGo/api/machine/v1"
)

func (c *ControllerV1) QueryCpus(ctx context.Context, req *v1.QueryCpusReq) (res *v1.QueryCpusRes, err error) {

	res = &v1.QueryCpusRes{}

	cpus, err := cpu.Info()

	physical, err := cpu.Counts(false)
	logical, err := cpu.Counts(true)

	res.PhysicalCore = physical
	res.LogicalCore = logical

	_ = gconv.Struct(cpus, &res.CPUS)

	return res, nil
}

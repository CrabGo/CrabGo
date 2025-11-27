package machine

import (
	"context"

	"github.com/shirou/gopsutil/mem"

	"CrabGo/api/machine/v1"
)

func (c *ControllerV1) QueryMemory(ctx context.Context, req *v1.QueryMemoryReq) (res *v1.QueryMemoryRes, err error) {

	vm, _ := mem.VirtualMemory()

	res = &v1.QueryMemoryRes{
		Total:       vm.Total / 1024 / 1024,
		Available:   vm.Available / 1024 / 1024,
		Used:        vm.Used / 1024 / 1024,
		UsedPercent: vm.UsedPercent,
		Free:        vm.Free / 1024 / 1024,
	}

	return res, nil
}

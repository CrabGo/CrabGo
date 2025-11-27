package machine

import (
	"CrabGo/internal/model"
	"CrabGo/internal/service"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/cpu"
)

type sMachine struct{}

func New() *sMachine {
	return &sMachine{}
}

func init() {
	service.RegisterMachine(New())
}

func (s *sMachine) QueryCpuInfo(ctx context.Context, input *model.QueryCpuInfoInput) (output *model.QueryCpuInfoOutput, err error) {

	//TODO:暂时没有查询条件

	cpus, err := cpu.Info()

	physical, err := cpu.Counts(false)
	logical, err := cpu.Counts(true)

	g.Log().Infof(ctx, "物理核:%v,逻辑核:%v", physical, logical)

	precentValue, err := cpu.Percent(time.Second, true)

	for _, v := range precentValue {

		g.Log().Infof(ctx, "percpu-true :%v", v)
	}

	precentValue, err = cpu.Percent(time.Second, false)

	for _, v := range precentValue {

		g.Log().Infof(ctx, "percpu-false:%v", v)
	}

	if err != nil {
		return nil, err
	}

	output = &model.QueryCpuInfoOutput{}

	err = gconv.Struct(cpus, &output)

	if err != nil {
		return nil, err
	}

	return output, nil
}

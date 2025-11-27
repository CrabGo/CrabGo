package machine

import (
	"CrabGo/api/machine/v1"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/process"
)

func (c *ControllerV1) QueryProcess(ctx context.Context, req *v1.QueryProcessReq) (res *v1.QueryProcessRes, err error) {

	res = &v1.QueryProcessRes{}

	var procesInfos []v1.ProcessInfoViewModel

	psal, err := process.Processes()

	for _, ps := range psal {

		isRunning, _ := ps.IsRunning()

		if !isRunning {
			continue
		}

		name, _ := ps.Name()
		cpu, _ := ps.CPUPercent()
		memory, _ := ps.MemoryPercent()
		threads, _ := ps.NumThreads()

		path, _ := ps.Cwd()
		exe, _ := ps.Exe()
		cmdLine, _ := ps.Cmdline()

		createTime, _ := ps.CreateTime()

		procesInfos = append(procesInfos, v1.ProcessInfoViewModel{
			Pid:        ps.Pid,
			Name:       name,
			CPU:        cpu,
			Memory:     memory,
			Threads:    threads,
			Path:       path,
			Exe:        exe,
			CmdLine:    cmdLine,
			CreateTime: createTime,
		})
	}

	gconv.Struct(procesInfos, &res)

	return res, nil
}

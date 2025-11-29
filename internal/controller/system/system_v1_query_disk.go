package system

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v4/disk"

	"CrabGo/api/system/v1"
)

func (c *ControllerV1) QueryDisk(ctx context.Context, req *v1.QueryDiskReq) (res *v1.QueryDiskRes, err error) {

	diskParts, err := disk.Partitions(false)

	if err != nil {
		return nil, err
	}

	res = &v1.QueryDiskRes{}

	var diskInfos []v1.DiskInfoViewModel

	for _, diskPart := range diskParts {

		diskUsed, err := disk.Usage(diskPart.Mountpoint)

		if err != nil {
			continue
		}

		diskInfos = append(diskInfos, v1.DiskInfoViewModel{
			Device:      diskPart.Device,
			MountPoint:  diskPart.Mountpoint,
			FsType:      diskPart.Fstype,
			Path:        diskUsed.Path,
			Total:       diskUsed.Total / 1024 / 1024,
			Free:        diskUsed.Free / 1024 / 1024,
			Used:        diskUsed.Used / 1024 / 1024,
			UsedPercent: diskUsed.UsedPercent,
		})
	}

	gconv.Struct(diskInfos, &res)

	return res, nil
}

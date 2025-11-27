package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QueryCpusReq 查询CPU信息
type QueryCpusReq struct {
	g.Meta `path:"/machine/cpu" tags:"machine" method:"get" summary:"查询CPU信息"`
}

// CpuInfoViewModel CPU信息
type CpuInfoViewModel struct {
	CPU        int32   `json:"cpu" dc:"CPU序号"`
	VendorID   string  `json:"vendorId" dc:"厂商"`
	Family     string  `json:"family" dc:""`
	Model      string  `json:"model" dc:"型号"`
	PhysicalID string  `json:"physicalId" dc:"物理ID"`
	CoreID     string  `json:"coreId" dc:"内核ID"`
	Cores      int32   `json:"cores" dc:"核心数"`
	ModelName  string  `json:"modelName" dc:"型号名称"`
	Mhz        float64 `json:"mhz" dc:"赫兹"`
}

// QueryCpusRes 查询CPU信息响应
type QueryCpusRes struct {
	PhysicalCore int                `json:"physical_core"`
	LogicalCore  int                `json:"logical_core"`
	CPUS         []CpuInfoViewModel `json:"cpus"`
}

// QueryCpuRateReq 查询CPU利用率
type QueryCpuRateReq struct {
	g.Meta `path:"/machine/cpu/use" tags:"machine" method:"get" summary:"查询CPU信息"`
}

// QueryCpuRateRes 查询CPU使用率响应
type QueryCpuRateRes struct {
	Use      float64   `json:"use" dc:"总体利用率"`
	CoreUser []float64 `json:"core_use" dc:"核心利用率"`
}

// QueryMemoryReq 查询内存请求
type QueryMemoryReq struct {
	g.Meta `path:"/machine/memory" tags:"machine" method:"get" summary:"查询内存信息"`
}

// QueryMemoryRes 查询内存响应
type QueryMemoryRes struct {
	Total       uint64  `json:"total" dc:"总内存"`
	Available   uint64  `json:"available" dc:"可用内存"`
	Used        uint64  `json:"used" dc:"已用内存"`
	UsedPercent float64 `json:"usedPercent" dc:"已使用百分比"`
	Free        uint64  `json:"free" dc:"剩余内存"`
}

// QueryDiskReq 查询磁盘
type QueryDiskReq struct {
	g.Meta `path:"/machine/disk" tags:"machine" method:"get" summary:"查询磁盘"`
}

// DiskInfoViewModel 磁盘信息
type DiskInfoViewModel struct {
	Device      string  `json:"device" dc:"设备"`
	MountPoint  string  `json:"mount_point"`
	FsType      string  `json:"fs_type"`
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

// QueryDiskRes 查询磁盘
type QueryDiskRes []DiskInfoViewModel

// QueryHostReq 查询主机
type QueryHostReq struct {
	g.Meta `path:"/machine/host" tags:"machine" method:"get" summary:"查询主机"`
}

// QueryHostRes 查询主机
type QueryHostRes struct {
	Hostname        string `json:"hostname"`
	Uptime          uint64 `json:"uptime"`
	BootTime        uint64 `json:"bootTime"`
	Procs           uint64 `json:"procs"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platformFamily"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	KernelArch      string `json:"kernelArch"`
	HostID          string `json:"hostId"`
}

type QueryNetReq struct {
	g.Meta `path:"/machine/net" tags:"machine" method:"get" summary:"查询主机"`
}

type NettInfoViewModel struct {
	Name    string `json:"name" dc:"网卡名称"`
	Mac     string `json:"mac" dc:"MAC地址"`
	IP4Addr string `json:"ip4"`
	IP6Addr string `json:"ip6"`
}
type QueryNetRes []NettInfoViewModel

type QueryProcessReq struct {
	g.Meta `path:"/machine/process" tags:"machine" method:"get" summary:"查询进程"`
}

type ProcessInfoViewModel struct {
	Pid        int32   `json:"pid"`
	Name       string  `json:"name"`
	CPU        float64 `json:"cpu"`
	Memory     float32 `json:"memory"`
	Threads    int32   `json:"threads"`
	Path       string  `json:"path"`
	Exe        string  `json:"exe"`
	CmdLine    string  `json:"cmd_line"`
	CreateTime int64   `json:"createTime"`
}
type QueryProcessRes []ProcessInfoViewModel

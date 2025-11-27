package machine

import (
	"context"
	"strings"

	"CrabGo/api/machine/v1"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/net"
)

func (c *ControllerV1) QueryNet(ctx context.Context, req *v1.QueryNetReq) (res *v1.QueryNetRes, err error) {
	is, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	res = &v1.QueryNetRes{}

	var netInfos []v1.NettInfoViewModel

	for _, ifi := range is {

		if ifi.Flags == nil || len(ifi.Flags) == 0 || len(ifi.Addrs) == 0 {
			continue
		}

		var flags = strings.Join(ifi.Flags, ",")

		if strings.Contains(flags, "loopback") || strings.Contains(ifi.Name, "docker") || strings.Contains(ifi.Name, "br-") {
			continue
		}

		var ip4addr string

		var ip6addr string

		for _, addr := range ifi.Addrs {
			if strings.Contains(addr.String(), "::") {
				ip6addr = addr.Addr
			} else {
				ip4addr = addr.Addr
			}
		}

		if ip4addr == "" {
			continue
		}

		netInfos = append(netInfos, v1.NettInfoViewModel{
			Name:    ifi.Name,
			Mac:     ifi.HardwareAddr,
			IP4Addr: ip4addr,
			IP6Addr: ip6addr,
		})
	}

	gconv.Struct(netInfos, &res)

	return res, nil
}

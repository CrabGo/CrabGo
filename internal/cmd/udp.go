package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
)

var UDP = &gcmd.Command{
	Name:        "udp",
	Brief:       "start udp server",
	Description: "this is the command entry for starting your grpc server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		fmt.Println("start grpc server")
		return
	},
}

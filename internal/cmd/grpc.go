package cmd

import (
	"CrabGo/internal/options/grpc"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Grpc = &gcmd.Command{
	Name:        "grpc",
	Brief:       "start grpc server",
	Description: "this is the command entry for starting your grpc server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

		log := g.Log("grpc")

		options, err := grpc.GetOptions(ctx)

		log.Infof(ctx, "GRPC服务监听地址%v", options.Address)

		go func() {
			<-ctx.Done()
			log.Info(ctx, "正在关闭Grpc服务")

		}()

		log.Info(ctx, "GRPC服务启动成功")

		return nil
	},
}

package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"CrabGo/internal/controller/hello"
)

var (
	API = &gcmd.Command{
		Name:  "api",
		Usage: "api",
		Brief: "start http api server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})

			go func() {
				<-ctx.Done()
				g.Log().Info(ctx, "正在关闭API服务")
				s.Shutdown()
			}()
			s.Run()

			return nil
		},
	}
)

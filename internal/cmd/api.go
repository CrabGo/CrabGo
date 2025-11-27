package cmd

import (
	"CrabGo/internal/controller/auth"
	"CrabGo/internal/controller/build"
	"CrabGo/internal/controller/machine"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
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
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					auth.NewV1(),
					machine.NewV1(),
					build.NewV1(),
				)
			})

			go func() {
				<-ctx.Done()
				g.Log().Info(ctx, "正在关闭API服务")
				s.Shutdown()
			}()

			s.BindStatusHandlerByMap(GetStatusHandlerFunc())
			s.SetGraceful(true)
			s.EnableAdmin()
			s.Run()

			return nil
		},
	}
)

func GetStatusHandlerFunc() map[int]ghttp.HandlerFunc {

	//200 OK：请求成功，返回请求的数据。
	//201 Created：资源创建成功，通常用于 POST 请求。
	//204 No Content：请求成功，但没有返回内容，通常用于 DELETE 请求。
	//400 Bad Request：请求参数错误或格式不正确
	//401 Unauthorized：未授权，需要身份验证。
	//403 Forbidden：已授权，但没有权限访问资源。
	//404 Not Found：请求的资源不存在。
	//500 Internal Server Error：服务器内部错误。

	handlerMap := make(map[int]ghttp.HandlerFunc)

	handlerMap[400] = func(r *ghttp.Request) {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeInvalidParameter.Code(),
			Message: "请求参数错误",
			Data:    nil,
		})
	}
	handlerMap[404] = func(r *ghttp.Request) {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotFound.Code(),
			Message: "请求终接点不存在",
			Data:    nil,
		})
	}
	handlerMap[500] = func(r *ghttp.Request) {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: "服务器内部错误,请稍后重试",
			Data:    nil,
		})
	}
	return handlerMap
}

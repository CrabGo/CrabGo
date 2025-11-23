package cmd

import (
	"CrabGo/internal/consts"
	"CrabGo/internal/model"
	"CrabGo/internal/options/tcp"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var TCP = &gcmd.Command{
	Name:        "tcp",
	Brief:       "start tcp server",
	Description: "this is the command entry for starting your grpc server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

		log := g.Log("tcp")

		options, err := tcp.GetOptions(ctx)

		log.Infof(ctx, "TCP服务监听地址%v", options.Address)

		server := g.TCPServer("tcp")
		server.SetAddress(options.Address)

		server.SetHandler(func(conn *gtcp.Conn) {
			defer func(conn *gtcp.Conn) {
				err := conn.Close()
				if err != nil {
					log.Errorf(ctx, "客户端关闭出现错误,错误信息:%v", err)
				}
			}(conn)

			for {

				data, err := conn.Recv(-1)

				if err != nil {
					continue
				}

				if len(data) <= 0 {
					continue
				}

				packet := model.NewPacket(data)

				err = packet.IsValid()

				if err != nil {
					log.Errorf(ctx, "客户端:%v 报文验证错误,错误信息:%v", conn.RemoteAddr(), err)
					continue
				}

				err = packet.Deserialize()

				if err != nil {
					log.Errorf(ctx, "客户端:%v 报文解析错误,错误信息:%v", conn.RemoteAddr(), err)
				}

				if packet.Type == consts.Undefined {
					continue
				}

			}
		})

		go func() {
			<-ctx.Done()
			log.Info(ctx, "正在关闭TCP服务")
			err := server.Close()
			if err != nil {
				log.Errorf(ctx, "服务端关闭出现错误,错误信息:%v", err)
				return
			}
		}()

		err = server.Run()

		log.Info(ctx, "TCP服务启动成功")

		if err != nil {
			log.Errorf(ctx, "TCP服务启动失败,%v", err)
		}

		return nil
	},
}

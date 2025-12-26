package cmd

import (
	"CrabGo/internal/options/udp"
	"context"
	"net"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gudp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var UDP = &gcmd.Command{
	Name:        "udp",
	Brief:       "start udp server",
	Description: "this is the command entry for starting your grpc server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

		log := g.Log("udp")

		options, err := udp.GetOptions(ctx)

		log.Infof(ctx, "UDP服务监听地址%v", options.Address)

		server := gudp.NewServer(options.Address, func(conn *gudp.ServerConn) {

			defer func() {
				err = conn.Close()
				if err != nil {
					log.Error(ctx, "关闭连接出现错误", err)
				}
			}()

			for {

				var addr *net.UDPAddr
				var data []byte
				data, addr, err = conn.Recv(-1)

				if err != nil {
					continue
				}

				if len(data) <= 0 {

					continue
				}

				log.Info(ctx, addr.IP.String())

				if len(data) == 0 {
				}
			}
		})
		go func() {
			<-ctx.Done()
			g.Log().Info(ctx, "正在关闭UDP服务")
			err := server.Close()
			if err != nil {
				log.Errorf(ctx, "服务端关闭出现错误,错误信息:%v", err)
				return
			}
		}()

		log.Info(ctx, "UDP服务启动成功")

		server.Run()

		return nil
	},
}

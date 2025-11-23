package cmd

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Main = &gcmd.Command{
	Description: "默认启动所有服务",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return ALL.Func(ctx, parser)
	},
}

var Help = &gcmd.Command{
	Name:  "help",
	Usage: "",
	Brief: "查看帮助",
	Description: `
		命令提示符
		---------------------------------------------------------------------------------
		启动服务
		>> 所有服务  [go run main.go]   热编译  [gf run main.go]
		>> API服务  [go run main.go api]
		>> TCP服务  [go run main.go tcp]
		>> UDP服务  [go run main.go udp]
		>> Grpc服务  [go run main.go grpc]
		>> 查看帮助  [go run main.go help]
		---------------------------------------------------------------------------------
    `,
	Arguments:     nil,
	Func:          nil,
	FuncWithValue: nil,
	HelpFunc:      nil,
	Examples:      "",
	Additional:    "",
	Strict:        false,
	CaseSensitive: false,
	Config:        "",
}

var ALL = &gcmd.Command{
	Name:        "all",
	Brief:       "启动所有服务",
	Description: "启动所有服务",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

		g.Log().Info(ctx, "正在启动所有服务")

		ctx, cancel := context.WithCancel(ctx)

		var wg sync.WaitGroup

		var allCommands = []*gcmd.Command{API, TCP, UDP, Grpc}

		for _, command := range allCommands {
			wg.Add(1)
			go func(cmd *gcmd.Command) {
				defer wg.Done()
				err := command.Func(ctx, parser)
				if err != nil {
					g.Log().Fatalf(ctx, "%v 启动出现错误:%v", command.Name, err)
				}

				g.Log().Infof(ctx, "%v 启动成功", command.Name)
			}(command)
		}
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			sig := <-sigs
			g.Log().Info(ctx, "收到信号:", sig)
			cancel()
			wg.Wait()
			os.Exit(0)
		}()
		select {}
		//g.Log().Info(ctx, "所有服务完成")

		return
	},
}

func init() {
	err := Main.AddCommand(ALL, API, TCP, UDP, Grpc, Help)
	if err != nil {
		panic(err)
	}
}

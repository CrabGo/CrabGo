package main

import (
	_ "CrabGo/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"CrabGo/internal/cmd"
)

func main() {
	//cmd.API.Run(gctx.GetInitCtx())

	err := cmd.Main.AddCommand(cmd.API, cmd.TCP, cmd.UDP, cmd.Grpc)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}

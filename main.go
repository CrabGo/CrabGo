package main

import (
	_ "CrabGo/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"CrabGo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

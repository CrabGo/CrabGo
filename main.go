package main

import (
	_ "CrabGo/internal/packed"
	"errors"

	_ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"CrabGo/internal/cmd"
)

func main() {

	var ctx = gctx.GetInitCtx()

	g.Log().Info(ctx, "正在检查数据库连接")

	err := connectDb()

	if err != nil {
		g.Log().Error(ctx, "数据库连接失败", err)
		panic(err)
	}

	g.Log().Info(ctx, "数据库连接成功")
	cmd.Main.Run(ctx)
}

func connectDb() error {
	err := g.DB().PingMaster()

	if err != nil {
		return errors.New("数据库连接失败")
	}

	return nil
}

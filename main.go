package main

import (
	_ "CrabGo/internal/logic"
	_ "CrabGo/internal/packed"
	"context"
	"errors"

	"CrabGo/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
)

func main() {

	var ctx = gctx.GetInitCtx()

	g.Log().Info(ctx, "正在检查配置")
	unPack(ctx)

	g.Log().Info(ctx, "正在检查数据库连接")
	err := connectDb()

	if err != nil {
		g.Log().Error(ctx, "数据库连接失败", err)
		panic(err)
	} else {
		g.Log().Info(ctx, "数据库连接成功")
	}

	cmd.Main.Run(ctx)
}

func unPack(ctx context.Context) {

	//参考配置文件中的资源文件夹 resource,manifest
	unPackResDir := g.SliceStr{"resource", "manifest"}
	for _, path := range unPackResDir {
		unPackResource(ctx, path)
	}
}

func unPackResource(ctx context.Context, dir string) {

	if gfile.Exists(dir) {

		return
	}

	g.Log().Errorf(ctx, "当前不存在目录%v正在解压初始配置", dir)
	files := gres.ScanDir(dir, "*.*", true)

	if len(files) == 0 {
		g.Log().Info(ctx, "%v下没有需要解压的资源")
		gfile.Mkdir(dir)
		return
	}
	var (
		err  error
		name string
		path string
	)

	for _, file := range files {
		name = file.Name()
		name = gstr.Trim(name, `\/`)
		if name == "" {
			continue
		}
		path = name
		if gfile.Exists(path) {
			continue
		}
		if file.FileInfo().IsDir() {
			err = gfile.Mkdir(path)
		} else {
			err = gfile.PutBytes(path, file.Content())
		}

		if err != nil {
			g.Log().Errorf(ctx, "解压资源%v失败,错误信息%v", path, err)
		} else {
			g.Log().Infof(ctx, "资源%v解压成功", name)
		}
	}
}

func connectDb() error {

	err := g.DB().PingMaster()

	if err != nil {

		link := g.DB().GetConfig().Link

		return errors.New("数据库连接失败,数据库地址:" + link)
	}

	return nil
}

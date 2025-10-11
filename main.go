package main

import (
	_ "CrabGo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"CrabGo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

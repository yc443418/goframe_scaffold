package main

import (
	"SmartRental/global"
	_ "SmartRental/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"SmartRental/internal/cmd"
)

func main() {

	var ctx = gctx.GetInitCtx()

	global.Init(ctx)

	cmd.Main.Run(ctx)
}

package main

import (
	_ "jiandazi-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"jiandazi-go/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2" // 导入pgsql包，以便在代码中使用pgsql数据库
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

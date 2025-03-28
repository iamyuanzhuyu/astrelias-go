package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.New()

	// 初始化数据库
	InitDatabase()

	// 执行数据库迁移
	// InitMigration()

	// 其他初始化逻辑...
	// 例如：InitCache(), InitLogger() 等

	g.Log().Info(ctx, "应用初始化完成")
}

package boot

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// InitDatabase 初始化数据库连接
func InitDatabase() {
	ctx := gctx.New()
	
	// 数据库初始化检查
	if g.DB().PingMaster() != nil {
		g.Log().Fatal(ctx, "数据库连接失败")
	}
	
	g.Log().Info(ctx, "数据库连接成功")
	
	// 其他数据库相关的初始化逻辑...
}
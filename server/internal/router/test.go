package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	testcheck "hotgo/internal/controller/api/test"
)

func Test(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		// 允许通过根地址访问的路由可以加到这里，访问地址：http://127.0.0.1:8000
		group.Bind(
			testcheck.TestCheck,
		)

	})
}

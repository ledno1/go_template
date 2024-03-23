// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package genrouter

import "hotgo/addons/clientupdate/controller/admin/sys"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, sys.ClientUpdate) // 插件客户端更新
}

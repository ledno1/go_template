// Package clientupdate
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package clientupdate

import (
	"hotgo/addons/clientupdate/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询插件客户端更新列表
type ListReq struct {
	g.Meta `path:"/clientUpdate/list" method:"get" tags:"插件客户端更新" summary:"获取插件客户端更新列表"`
	sysin.ClientUpdateListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.ClientUpdateListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取插件客户端更新指定信息
type ViewReq struct {
	g.Meta `path:"/clientUpdate/view" method:"get" tags:"插件客户端更新" summary:"获取插件客户端更新指定信息"`
	sysin.ClientUpdateViewInp
}

type ViewRes struct {
	*sysin.ClientUpdateViewModel
}

// EditReq 修改/新增插件客户端更新
type EditReq struct {
	g.Meta `path:"/clientUpdate/edit" method:"post" tags:"插件客户端更新" summary:"修改/新增插件客户端更新"`
	sysin.ClientUpdateEditInp
}

type EditRes struct{}

// DeleteReq 删除插件客户端更新
type DeleteReq struct {
	g.Meta `path:"/clientUpdate/delete" method:"post" tags:"插件客户端更新" summary:"删除插件客户端更新"`
	sysin.ClientUpdateDeleteInp
}

type DeleteRes struct{}

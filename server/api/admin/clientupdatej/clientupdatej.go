// Package clientupdatej
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package clientupdatej

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询插件客户端更新列表
type ListReq struct {
	g.Meta `path:"/clientUpdateJ/list" method:"get" tags:"插件客户端更新" summary:"获取插件客户端更新列表"`
	sysin.ClientUpdateJListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.ClientUpdateJListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出插件客户端更新列表
type ExportReq struct {
	g.Meta `path:"/clientUpdateJ/export" method:"get" tags:"插件客户端更新" summary:"导出插件客户端更新列表"`
	sysin.ClientUpdateJListInp
}

type ExportRes struct{}

// ViewReq 获取插件客户端更新指定信息
type ViewReq struct {
	g.Meta `path:"/clientUpdateJ/view" method:"get" tags:"插件客户端更新" summary:"获取插件客户端更新指定信息"`
	sysin.ClientUpdateJViewInp
}

type ViewRes struct {
	*sysin.ClientUpdateJViewModel
}

// EditReq 修改/新增插件客户端更新
type EditReq struct {
	g.Meta `path:"/clientUpdateJ/edit" method:"post" tags:"插件客户端更新" summary:"修改/新增插件客户端更新"`
	sysin.ClientUpdateJEditInp
}

type EditRes struct{}

// DeleteReq 删除插件客户端更新
type DeleteReq struct {
	g.Meta `path:"/clientUpdateJ/delete" method:"post" tags:"插件客户端更新" summary:"删除插件客户端更新"`
	sysin.ClientUpdateJDeleteInp
}

type DeleteRes struct{}

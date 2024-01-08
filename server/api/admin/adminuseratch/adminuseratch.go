// Package adminuseratch
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package adminuseratch

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询授权/密钥列表
type ListReq struct {
	g.Meta `path:"/adminUserAtch/list" method:"get" tags:"授权/密钥" summary:"获取授权/密钥列表"`
	sysin.AdminUserAtchListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AdminUserAtchListModel `json:"list"   dc:"数据列表"`
}

// ViewReq 获取授权/密钥指定信息
type ViewReq struct {
	g.Meta `path:"/adminUserAtch/view" method:"get" tags:"授权/密钥" summary:"获取授权/密钥指定信息"`
	sysin.AdminUserAtchViewInp
}

type ViewRes struct {
	*sysin.AdminUserAtchViewModel
}

// EditReq 修改/新增授权/密钥
type EditReq struct {
	g.Meta `path:"/adminUserAtch/edit" method:"post" tags:"授权/密钥" summary:"修改/新增授权/密钥"`
	sysin.AdminUserAtchEditInp
}

type EditRes struct{}

// DeleteReq 删除授权/密钥
type DeleteReq struct {
	g.Meta `path:"/adminUserAtch/delete" method:"post" tags:"授权/密钥" summary:"删除授权/密钥"`
	sysin.AdminUserAtchDeleteInp
}

type DeleteRes struct{}

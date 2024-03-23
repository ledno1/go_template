// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/clientupdate/api/admin/clientupdate"
	"hotgo/addons/clientupdate/service"
)

var (
	ClientUpdate = cClientUpdate{}
)

type cClientUpdate struct{}

// List 查看插件客户端更新列表
func (c *cClientUpdate) List(ctx context.Context, req *clientupdate.ListReq) (res *clientupdate.ListRes, err error) {
	list, totalCount, err := service.SysClientUpdate().List(ctx, &req.ClientUpdateListInp)
	if err != nil {
		return
	}

	res = new(clientupdate.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新插件客户端更新
func (c *cClientUpdate) Edit(ctx context.Context, req *clientupdate.EditReq) (res *clientupdate.EditRes, err error) {
	err = service.SysClientUpdate().Edit(ctx, &req.ClientUpdateEditInp)
	return
}

// View 获取指定插件客户端更新信息
func (c *cClientUpdate) View(ctx context.Context, req *clientupdate.ViewReq) (res *clientupdate.ViewRes, err error) {
	data, err := service.SysClientUpdate().View(ctx, &req.ClientUpdateViewInp)
	if err != nil {
		return
	}

	res = new(clientupdate.ViewRes)
	res.ClientUpdateViewModel = data
	return
}

// Delete 删除插件客户端更新
func (c *cClientUpdate) Delete(ctx context.Context, req *clientupdate.DeleteReq) (res *clientupdate.DeleteRes, err error) {
	err = service.SysClientUpdate().Delete(ctx, &req.ClientUpdateDeleteInp)
	return
}

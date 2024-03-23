// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/api/admin/clientupdatej"
	"hotgo/internal/service"
)

var (
	ClientUpdateJ = cClientUpdateJ{}
)

type cClientUpdateJ struct{}

// List 查看插件客户端更新列表
func (c *cClientUpdateJ) List(ctx context.Context, req *clientupdatej.ListReq) (res *clientupdatej.ListRes, err error) {
	list, totalCount, err := service.SysClientUpdateJ().List(ctx, &req.ClientUpdateJListInp)
	if err != nil {
		return
	}

	res = new(clientupdatej.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出插件客户端更新列表
func (c *cClientUpdateJ) Export(ctx context.Context, req *clientupdatej.ExportReq) (res *clientupdatej.ExportRes, err error) {
	err = service.SysClientUpdateJ().Export(ctx, &req.ClientUpdateJListInp)
	return
}

// Edit 更新插件客户端更新
func (c *cClientUpdateJ) Edit(ctx context.Context, req *clientupdatej.EditReq) (res *clientupdatej.EditRes, err error) {
	err = service.SysClientUpdateJ().Edit(ctx, &req.ClientUpdateJEditInp)
	return
}

// View 获取指定插件客户端更新信息
func (c *cClientUpdateJ) View(ctx context.Context, req *clientupdatej.ViewReq) (res *clientupdatej.ViewRes, err error) {
	data, err := service.SysClientUpdateJ().View(ctx, &req.ClientUpdateJViewInp)
	if err != nil {
		return
	}

	res = new(clientupdatej.ViewRes)
	res.ClientUpdateJViewModel = data
	return
}

// Delete 删除插件客户端更新
func (c *cClientUpdateJ) Delete(ctx context.Context, req *clientupdatej.DeleteReq) (res *clientupdatej.DeleteRes, err error) {
	err = service.SysClientUpdateJ().Delete(ctx, &req.ClientUpdateJDeleteInp)
	return
}

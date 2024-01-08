// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/api/admin/adminuseratch"
	"hotgo/internal/service"
)

var (
	AdminUserAtch = cAdminUserAtch{}
)

type cAdminUserAtch struct{}

// List 查看授权/密钥列表
func (c *cAdminUserAtch) List(ctx context.Context, req *adminuseratch.ListReq) (res *adminuseratch.ListRes, err error) {
	list, totalCount, err := service.SysAdminUserAtch().List(ctx, &req.AdminUserAtchListInp)
	if err != nil {
		return
	}

	res = new(adminuseratch.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Edit 更新授权/密钥
func (c *cAdminUserAtch) Edit(ctx context.Context, req *adminuseratch.EditReq) (res *adminuseratch.EditRes, err error) {
	err = service.SysAdminUserAtch().Edit(ctx, &req.AdminUserAtchEditInp)
	return
}

// View 获取指定授权/密钥信息
func (c *cAdminUserAtch) View(ctx context.Context, req *adminuseratch.ViewReq) (res *adminuseratch.ViewRes, err error) {
	data, err := service.SysAdminUserAtch().View(ctx, &req.AdminUserAtchViewInp)
	if err != nil {
		return
	}

	res = new(adminuseratch.ViewRes)
	res.AdminUserAtchViewModel = data
	return
}

// Delete 删除授权/密钥
func (c *cAdminUserAtch) Delete(ctx context.Context, req *adminuseratch.DeleteReq) (res *adminuseratch.DeleteRes, err error) {
	err = service.SysAdminUserAtch().Delete(ctx, &req.AdminUserAtchDeleteInp)
	return
}

// Package api
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package api

import (
	"context"
	"hotgo/addons/clientupdate/api/api/index"
	"hotgo/addons/clientupdate/service"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

// Test 测试
func (c *cIndex) Test(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error) {
	data, err := service.SysIndex().Test(ctx, &req.IndexTestInp)
	if err != nil {
		return
	}
	res = new(index.TestRes)
	res.IndexTestModel = data
	return
}

// GetUpdate 获取是否需要更新和更新下载地址
func (c *cIndex) GetUpdate(ctx context.Context, req *index.GetUpdateReq) (res *index.GetUpdateRes, err error) {
	// 判断逻辑
	// 1判断是否有该appName的更新记录 无则判断无需更新
	// 2 判断版本是否相同 不相同则返回最新(创建时间最晚)版本地址
	//
	return
}

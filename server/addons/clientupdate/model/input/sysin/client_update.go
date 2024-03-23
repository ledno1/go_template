// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientUpdateUpdateFields 修改插件客户端更新字段过滤
type ClientUpdateUpdateFields struct {
	AppName    int    `json:"appName"    dc:"应用类型"`
	AppVersion string `json:"appVersion" dc:"应用版本"`
	AppUrl     string `json:"appUrl"     dc:"更新包地址"`
	Mark       string `json:"mark"       dc:"备忘"`
}

// ClientUpdateInsertFields 新增插件客户端更新字段过滤
type ClientUpdateInsertFields struct {
	AppName    int    `json:"appName"    dc:"应用类型"`
	AppVersion string `json:"appVersion" dc:"应用版本"`
	AppUrl     string `json:"appUrl"     dc:"更新包地址"`
	Mark       string `json:"mark"       dc:"备忘"`
}

// ClientUpdateEditInp 修改/新增插件客户端更新
type ClientUpdateEditInp struct {
	entity.AddonClientUpdate
}

func (in *ClientUpdateEditInp) Filter(ctx context.Context) (err error) {
	// 验证应用类型
	if err := g.Validator().Rules("required").Data(in.AppName).Messages("应用类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1").Data(in.AppName).Messages("应用类型值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证应用版本
	if err := g.Validator().Rules("required").Data(in.AppVersion).Messages("应用版本不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证更新包地址
	if err := g.Validator().Rules("required").Data(in.AppUrl).Messages("更新包地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type ClientUpdateEditModel struct{}

// ClientUpdateDeleteInp 删除插件客户端更新
type ClientUpdateDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ClientUpdateDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ClientUpdateDeleteModel struct{}

// ClientUpdateViewInp 获取指定插件客户端更新信息
type ClientUpdateViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ClientUpdateViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ClientUpdateViewModel struct {
	entity.AddonClientUpdate
}

// ClientUpdateListInp 获取插件客户端更新列表
type ClientUpdateListInp struct {
	form.PageReq
	AppName    int           `json:"appName"    dc:"应用类型"`
	AppVersion string        `json:"appVersion" dc:"应用版本"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"创建时间"`
}

func (in *ClientUpdateListInp) Filter(ctx context.Context) (err error) {
	return
}

type ClientUpdateListModel struct {
	Id         int         `json:"id"         dc:"id"`
	AppName    int         `json:"appName"    dc:"应用类型"`
	AppVersion string      `json:"appVersion" dc:"应用版本"`
	AppUrl     string      `json:"appUrl"     dc:"更新包地址"`
	Mark       string      `json:"mark"       dc:"备忘"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

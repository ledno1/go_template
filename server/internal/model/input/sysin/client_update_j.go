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

// ClientUpdateJUpdateFields 修改插件客户端更新字段过滤
type ClientUpdateJUpdateFields struct {
	AppName    int    `json:"appName"    dc:"应用类型"`
	AppVersion string `json:"appVersion" dc:"应用版本"`
	AppUrl     string `json:"appUrl"     dc:"更新包地址"`
	Mark       string `json:"mark"       dc:"备忘"`
}

// ClientUpdateJInsertFields 新增插件客户端更新字段过滤
type ClientUpdateJInsertFields struct {
	AppName    int    `json:"appName"    dc:"应用类型"`
	AppVersion string `json:"appVersion" dc:"应用版本"`
	AppUrl     string `json:"appUrl"     dc:"更新包地址"`
	Mark       string `json:"mark"       dc:"备忘"`
}

// ClientUpdateJEditInp 修改/新增插件客户端更新
type ClientUpdateJEditInp struct {
	entity.ClientUpdate
}

func (in *ClientUpdateJEditInp) Filter(ctx context.Context) (err error) {
	// 验证应用类型
	if err := g.Validator().Rules("required").Data(in.AppName).Messages("应用类型不能为空").Run(ctx); err != nil {
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

	// 验证备忘
	if err := g.Validator().Rules("required").Data(in.Mark).Messages("备忘不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type ClientUpdateJEditModel struct{}

// ClientUpdateJDeleteInp 删除插件客户端更新
type ClientUpdateJDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ClientUpdateJDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ClientUpdateJDeleteModel struct{}

// ClientUpdateJViewInp 获取指定插件客户端更新信息
type ClientUpdateJViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ClientUpdateJViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ClientUpdateJViewModel struct {
	entity.ClientUpdate
}

// ClientUpdateJListInp 获取插件客户端更新列表
type ClientUpdateJListInp struct {
	form.PageReq
	AppName   int           `json:"appName"   dc:"应用类型"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ClientUpdateJListInp) Filter(ctx context.Context) (err error) {
	return
}

type ClientUpdateJListModel struct {
	Id         int         `json:"id"         dc:"id"`
	AppName    int         `json:"appName"    dc:"应用类型"`
	AppVersion string      `json:"appVersion" dc:"应用版本"`
	Mark       string      `json:"mark"       dc:"备忘"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// ClientUpdateJExportModel 导出插件客户端更新
type ClientUpdateJExportModel struct {
	Id         int         `json:"id"         dc:"id"`
	AppName    int         `json:"appName"    dc:"应用类型"`
	AppVersion string      `json:"appVersion" dc:"应用版本"`
	Mark       string      `json:"mark"       dc:"备忘"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

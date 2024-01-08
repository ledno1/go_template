// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
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

// AdminUserAtchUpdateFields 修改授权/密钥字段过滤
type AdminUserAtchUpdateFields struct {
	Uid              int    `json:"uid"              dc:"管理员ID"`
	HaxKey           string `json:"haxKey"           dc:"密钥"`
	AuthCount        int    `json:"authCount"        dc:"授权总数"`
	AuthSuccessCount int    `json:"authSuccessCount" dc:"授权成功"`
}

// AdminUserAtchInsertFields 新增授权/密钥字段过滤
type AdminUserAtchInsertFields struct {
	Uid              int    `json:"uid"              dc:"管理员ID"`
	HaxKey           string `json:"haxKey"           dc:"密钥"`
	AuthCount        int    `json:"authCount"        dc:"授权总数"`
	AuthSuccessCount int    `json:"authSuccessCount" dc:"授权成功"`
}

// AdminUserAtchEditInp 修改/新增授权/密钥
type AdminUserAtchEditInp struct {
	entity.AdminUserAtch
}

func (in *AdminUserAtchEditInp) Filter(ctx context.Context) (err error) {
	// 验证管理员ID
	if err := g.Validator().Rules("required").Data(in.Uid).Messages("管理员ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证密钥
	if err := g.Validator().Rules("required").Data(in.HaxKey).Messages("密钥不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证授权总数
	if err := g.Validator().Rules("required").Data(in.AuthCount).Messages("授权总数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证授权成功
	if err := g.Validator().Rules("required").Data(in.AuthSuccessCount).Messages("授权成功不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type AdminUserAtchEditModel struct{}

// AdminUserAtchDeleteInp 删除授权/密钥
type AdminUserAtchDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *AdminUserAtchDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AdminUserAtchDeleteModel struct{}

// AdminUserAtchViewInp 获取指定授权/密钥信息
type AdminUserAtchViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *AdminUserAtchViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AdminUserAtchViewModel struct {
	entity.AdminUserAtch
}

// AdminUserAtchListInp 获取授权/密钥列表
type AdminUserAtchListInp struct {
	form.PageReq
	CreatedAt           []*gtime.Time `json:"createdAt"           dc:"创建时间"`
	AdminMemberRealName string        `json:"adminMemberRealName" dc:"真实姓名"`
	AdminMemberUsername string        `json:"adminMemberUsername" dc:"帐号"`
}

func (in *AdminUserAtchListInp) Filter(ctx context.Context) (err error) {
	return
}

type AdminUserAtchListModel struct {
	Id                  int         `json:"id"                  dc:"id"`
	Uid                 int         `json:"uid"                 dc:"管理员ID"`
	HaxKey              string      `json:"haxKey"              dc:"密钥"`
	AuthCount           int         `json:"authCount"           dc:"授权总数"`
	AuthSuccessCount    int         `json:"authSuccessCount"    dc:"授权成功"`
	CreatedAt           *gtime.Time `json:"createdAt"           dc:"创建时间"`
	UpdatedAt           *gtime.Time `json:"updatedAt"           dc:"更新时间"`
	AdminMemberRealName string      `json:"adminMemberRealName" dc:"真实姓名"`
	AdminMemberUsername string      `json:"adminMemberUsername" dc:"帐号"`
}

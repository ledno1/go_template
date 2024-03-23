// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/addons/clientupdate/model/input/sysin"
	"hotgo/addons/clientupdate/service"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sSysClientUpdate struct{}

func NewSysClientUpdate() *sSysClientUpdate {
	return &sSysClientUpdate{}
}

func init() {
	service.RegisterSysClientUpdate(NewSysClientUpdate())
}

// Model 插件客户端更新ORM模型
func (s *sSysClientUpdate) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AddonClientUpdate.Ctx(ctx), option...)
}

// List 获取插件客户端更新列表
func (s *sSysClientUpdate) List(ctx context.Context, in *sysin.ClientUpdateListInp) (list []*sysin.ClientUpdateListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询应用类型
	if in.AppName > 0 {
		mod = mod.Where(dao.AddonClientUpdate.Columns().AppName, in.AppName)
	}

	// 查询应用版本
	if in.AppVersion != "" {
		mod = mod.WhereLike(dao.AddonClientUpdate.Columns().AppVersion, "%"+in.AppVersion+"%")
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AddonClientUpdate.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取插件客户端更新数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.ClientUpdateListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.AddonClientUpdate.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取插件客户端更新列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增插件客户端更新
func (s *sSysClientUpdate) Edit(ctx context.Context, in *sysin.ClientUpdateEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.ClientUpdateUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改插件客户端更新失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.ClientUpdateInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增插件客户端更新失败，请稍后重试！")
	}
	return
}

// Delete 删除插件客户端更新
func (s *sSysClientUpdate) Delete(ctx context.Context, in *sysin.ClientUpdateDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除插件客户端更新失败，请稍后重试！")
		return
	}
	return
}

// View 获取插件客户端更新指定信息
func (s *sSysClientUpdate) View(ctx context.Context, in *sysin.ClientUpdateViewInp) (res *sysin.ClientUpdateViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取插件客户端更新信息，请稍后重试！")
		return
	}
	return
}

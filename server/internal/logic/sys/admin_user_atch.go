// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sSysAdminUserAtch struct{}

func NewSysAdminUserAtch() *sSysAdminUserAtch {
	return &sSysAdminUserAtch{}
}

func init() {
	service.RegisterSysAdminUserAtch(NewSysAdminUserAtch())
}

// Model 授权/密钥ORM模型
func (s *sSysAdminUserAtch) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AdminUserAtch.Ctx(ctx), option...)
}

// List 获取授权/密钥列表
func (s *sSysAdminUserAtch) List(ctx context.Context, in *sysin.AdminUserAtchListInp) (list []*sysin.AdminUserAtchListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AdminUserAtch.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 查询真实姓名
	if in.AdminMemberRealName != "" {
		mod = mod.WhereLike(dao.AdminMember.Columns().RealName, in.AdminMemberRealName)
	}

	// 查询帐号
	if in.AdminMemberUsername != "" {
		mod = mod.WhereLike(dao.AdminMember.Columns().Username, in.AdminMemberUsername)
	}

	// 关联表adminMember
	mod = mod.LeftJoin(hgorm.GenJoinOnRelation(
		dao.AdminUserAtch.Table(), dao.AdminUserAtch.Columns().Uid, // 主表表名,关联字段
		dao.AdminMember.Table(), "adminMember", dao.AdminMember.Columns().Id, // 关联表表名,别名,关联字段
	)...)

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取授权/密钥数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	// 关联表select
	fields, err := hgorm.GenJoinSelect(ctx, sysin.AdminUserAtchListModel{}, &dao.AdminUserAtch, []*hgorm.Join{
		{Dao: &dao.AdminMember, Alias: "adminMember"},
	})

	if err != nil {
		err = gerror.Wrap(err, "获取授权/密钥关联字段失败，请稍后重试！")
		return
	}

	if err = mod.Fields(fields).Page(in.Page, in.PerPage).OrderDesc(dao.AdminUserAtch.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取授权/密钥列表失败，请稍后重试！")
		return
	}
	return
}

// Edit 修改/新增授权/密钥
func (s *sSysAdminUserAtch) Edit(ctx context.Context, in *sysin.AdminUserAtchEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.AdminUserAtchUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改授权/密钥失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.AdminUserAtchInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增授权/密钥失败，请稍后重试！")
	}
	return
}

// Delete 删除授权/密钥
func (s *sSysAdminUserAtch) Delete(ctx context.Context, in *sysin.AdminUserAtchDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除授权/密钥失败，请稍后重试！")
		return
	}
	return
}

// View 获取授权/密钥指定信息
func (s *sSysAdminUserAtch) View(ctx context.Context, in *sysin.AdminUserAtchViewInp) (res *sysin.AdminUserAtchViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取授权/密钥信息，请稍后重试！")
		return
	}
	return
}

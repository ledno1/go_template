// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.9.3
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysClientUpdateJ struct{}

func NewSysClientUpdateJ() *sSysClientUpdateJ {
	return &sSysClientUpdateJ{}
}

func init() {
	service.RegisterSysClientUpdateJ(NewSysClientUpdateJ())
}

// Model 插件客户端更新ORM模型
func (s *sSysClientUpdateJ) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ClientUpdate.Ctx(ctx), option...)
}

// List 获取插件客户端更新列表
func (s *sSysClientUpdateJ) List(ctx context.Context, in *sysin.ClientUpdateJListInp) (list []*sysin.ClientUpdateJListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 查询应用类型
	if in.AppName > 0 {
		mod = mod.Where(dao.ClientUpdate.Columns().AppName, in.AppName)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.ClientUpdate.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	totalCount, err = mod.Clone().Count()
	if err != nil {
		err = gerror.Wrap(err, "获取插件客户端更新数据行失败，请稍后重试！")
		return
	}

	if totalCount == 0 {
		return
	}

	if err = mod.Fields(sysin.ClientUpdateJListModel{}).Page(in.Page, in.PerPage).OrderDesc(dao.ClientUpdate.Columns().Id).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取插件客户端更新列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出插件客户端更新
func (s *sSysClientUpdateJ) Export(ctx context.Context, in *sysin.ClientUpdateJListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.ClientUpdateJExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出插件客户端更新-" + gctx.CtxId(ctx) + ".xlsx"
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.ClientUpdateJExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增插件客户端更新
func (s *sSysClientUpdateJ) Edit(ctx context.Context, in *sysin.ClientUpdateJEditInp) (err error) {
	// 修改
	if in.Id > 0 {
		if _, err = s.Model(ctx).
			Fields(sysin.ClientUpdateJUpdateFields{}).
			WherePri(in.Id).Data(in).Update(); err != nil {
			err = gerror.Wrap(err, "修改插件客户端更新失败，请稍后重试！")
		}
		return
	}

	// 新增
	if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
		Fields(sysin.ClientUpdateJInsertFields{}).
		Data(in).Insert(); err != nil {
		err = gerror.Wrap(err, "新增插件客户端更新失败，请稍后重试！")
	}
	return
}

// Delete 删除插件客户端更新
func (s *sSysClientUpdateJ) Delete(ctx context.Context, in *sysin.ClientUpdateJDeleteInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Delete(); err != nil {
		err = gerror.Wrap(err, "删除插件客户端更新失败，请稍后重试！")
		return
	}
	return
}

// View 获取插件客户端更新指定信息
func (s *sSysClientUpdateJ) View(ctx context.Context, in *sysin.ClientUpdateJViewInp) (res *sysin.ClientUpdateJViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取插件客户端更新信息，请稍后重试！")
		return
	}
	return
}

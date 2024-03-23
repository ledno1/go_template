// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/clientupdate/model"
	"hotgo/addons/clientupdate/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysClientUpdate interface {
		// Model 插件客户端更新ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取插件客户端更新列表
		List(ctx context.Context, in *sysin.ClientUpdateListInp) (list []*sysin.ClientUpdateListModel, totalCount int, err error)
		// Edit 修改/新增插件客户端更新
		Edit(ctx context.Context, in *sysin.ClientUpdateEditInp) (err error)
		// Delete 删除插件客户端更新
		Delete(ctx context.Context, in *sysin.ClientUpdateDeleteInp) (err error)
		// View 获取插件客户端更新指定信息
		View(ctx context.Context, in *sysin.ClientUpdateViewInp) (res *sysin.ClientUpdateViewModel, err error)
	}
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
)

var (
	localSysConfig       ISysConfig
	localSysIndex        ISysIndex
	localSysClientUpdate ISysClientUpdate
)

func SysClientUpdate() ISysClientUpdate {
	if localSysClientUpdate == nil {
		panic("implement not found for interface ISysClientUpdate, forgot register?")
	}
	return localSysClientUpdate
}

func RegisterSysClientUpdate(i ISysClientUpdate) {
	localSysClientUpdate = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}

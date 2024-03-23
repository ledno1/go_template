// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddonClientUpdateDao is the data access object for table hg_addon_client_update.
type AddonClientUpdateDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns AddonClientUpdateColumns // columns contains all the column names of Table for convenient usage.
}

// AddonClientUpdateColumns defines and stores column names for table hg_addon_client_update.
type AddonClientUpdateColumns struct {
	Id         string //
	AppName    string // 应用类型
	AppVersion string // 应用版本
	AppUrl     string // 更新包地址
	Mark       string // 备忘
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// addonClientUpdateColumns holds the columns for table hg_addon_client_update.
var addonClientUpdateColumns = AddonClientUpdateColumns{
	Id:         "id",
	AppName:    "app_name",
	AppVersion: "app_version",
	AppUrl:     "app_url",
	Mark:       "mark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewAddonClientUpdateDao creates and returns a new DAO object for table data access.
func NewAddonClientUpdateDao() *AddonClientUpdateDao {
	return &AddonClientUpdateDao{
		group:   "default",
		table:   "hg_addon_client_update",
		columns: addonClientUpdateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AddonClientUpdateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AddonClientUpdateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AddonClientUpdateDao) Columns() AddonClientUpdateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AddonClientUpdateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AddonClientUpdateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AddonClientUpdateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

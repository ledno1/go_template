// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientUpdateDao is the data access object for table hg_client_update.
type ClientUpdateDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ClientUpdateColumns // columns contains all the column names of Table for convenient usage.
}

// ClientUpdateColumns defines and stores column names for table hg_client_update.
type ClientUpdateColumns struct {
	Id         string //
	AppName    string // 应用类型
	AppVersion string // 应用版本
	AppUrl     string // 更新包地址
	Mark       string // 备忘
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// clientUpdateColumns holds the columns for table hg_client_update.
var clientUpdateColumns = ClientUpdateColumns{
	Id:         "id",
	AppName:    "app_name",
	AppVersion: "app_version",
	AppUrl:     "app_url",
	Mark:       "mark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewClientUpdateDao creates and returns a new DAO object for table data access.
func NewClientUpdateDao() *ClientUpdateDao {
	return &ClientUpdateDao{
		group:   "default",
		table:   "hg_client_update",
		columns: clientUpdateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientUpdateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientUpdateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientUpdateDao) Columns() ClientUpdateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientUpdateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientUpdateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientUpdateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

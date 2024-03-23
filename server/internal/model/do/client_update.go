// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientUpdate is the golang structure of table hg_client_update for DAO operations like Where/Data.
type ClientUpdate struct {
	g.Meta     `orm:"table:hg_client_update, do:true"`
	Id         interface{} //
	AppName    interface{} // 应用类型
	AppVersion interface{} // 应用版本
	AppUrl     interface{} // 更新包地址
	Mark       interface{} // 备忘
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}

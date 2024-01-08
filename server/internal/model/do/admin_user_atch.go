// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUserAtch is the golang structure of table hg_admin_user_atch for DAO operations like Where/Data.
type AdminUserAtch struct {
	g.Meta           `orm:"table:hg_admin_user_atch, do:true"`
	Id               interface{} //
	Uid              interface{} // 用户id
	HaxKey           interface{} // 密钥
	AuthCount        interface{} // 授权总数
	AuthSuccessCount interface{} // 授权成功
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}

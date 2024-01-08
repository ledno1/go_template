// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUserAtch is the golang structure for table admin_user_atch.
type AdminUserAtch struct {
	Id               int         `json:"id"               description:""`
	Uid              int         `json:"uid"              description:"用户id"`
	HaxKey           string      `json:"haxKey"           description:"密钥"`
	AuthCount        int         `json:"authCount"        description:"授权总数"`
	AuthSuccessCount int         `json:"authSuccessCount" description:"授权成功"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
}

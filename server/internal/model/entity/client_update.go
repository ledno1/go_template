// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientUpdate is the golang structure for table client_update.
type ClientUpdate struct {
	Id         int         `json:"id"         description:""`
	AppName    int         `json:"appName"    description:"应用类型"`
	AppVersion string      `json:"appVersion" description:"应用版本"`
	AppUrl     string      `json:"appUrl"     description:"更新包地址"`
	Mark       string      `json:"mark"       description:"备忘"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
}

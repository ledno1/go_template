// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"hotgo/addons/clientupdate/api/admin"
)

type ControllerClientupdate struct{}

func NewClientupdate() admin.IAdminClientupdate {
	return &ControllerClientupdate{}
}

type ControllerConfig struct{}

func NewConfig() admin.IAdminConfig {
	return &ControllerConfig{}
}

type ControllerIndex struct{}

func NewIndex() admin.IAdminIndex {
	return &ControllerIndex{}
}


// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"hotgo/addons/clientupdate/api/admin/clientupdate"
	"hotgo/addons/clientupdate/api/admin/config"
	"hotgo/addons/clientupdate/api/admin/index"
)

type IAdminClientupdate interface {
	List(ctx context.Context, req *clientupdate.ListReq) (res *clientupdate.ListRes, err error)
	View(ctx context.Context, req *clientupdate.ViewReq) (res *clientupdate.ViewRes, err error)
	Edit(ctx context.Context, req *clientupdate.EditReq) (res *clientupdate.EditRes, err error)
	Delete(ctx context.Context, req *clientupdate.DeleteReq) (res *clientupdate.DeleteRes, err error)
}

type IAdminConfig interface {
	Get(ctx context.Context, req *config.GetReq) (res *config.GetRes, err error)
	Update(ctx context.Context, req *config.UpdateReq) (res *config.UpdateRes, err error)
}

type IAdminIndex interface {
	Test(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error)
}



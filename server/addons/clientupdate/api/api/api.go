// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package api

import (
	"context"
	
	"hotgo/addons/clientupdate/api/api/index"
)

type IApiIndex interface {
	Test(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error)
}



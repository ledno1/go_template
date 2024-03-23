// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package api

import (
	"context"
	
	"hotgo/api/api/member"
	"hotgo/api/api/pay"
	"hotgo/api/api/pcclient"
	"hotgo/api/api/test"
	"hotgo/api/api/user"
)

type IApiMember interface {
	GetIdByCode(ctx context.Context, req *member.GetIdByCodeReq) (res *member.GetIdByCodeRes, err error)
}

type IApiPay interface {
	NotifyAliPay(ctx context.Context, req *pay.NotifyAliPayReq) (res *pay.NotifyAliPayRes, err error)
	NotifyWxPay(ctx context.Context, req *pay.NotifyWxPayReq) (res *pay.NotifyWxPayRes, err error)
	NotifyQQPay(ctx context.Context, req *pay.NotifyQQPayReq) (res *pay.NotifyQQPayRes, err error)
}

type IApiPcclient interface {
	ClientUpdate(ctx context.Context, req *pcclient.ClientUpdateReq) (res *pcclient.ClientUpdateRes, err error)
}

type IApiTest interface {
	Add(ctx context.Context, req *test.AddReq) (res *test.AddRes, err error)
}

type IApiUser interface {
	Hello(ctx context.Context, req *user.HelloReq) (res *user.HelloRes, err error)
}



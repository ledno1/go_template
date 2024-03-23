package pcclient

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/api/api/pcclient"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
)

var (
	PcClient = cPcClient{}
)

type cPcClient struct{}

func (c *cPcClient) ClientUpdate(ctx context.Context, req *pcclient.ClientUpdateReq) (res *pcclient.ClientUpdateRes, err error) {
	g.Log().Debug(ctx, "更新询问")
	// 判断逻辑
	res = new(pcclient.ClientUpdateRes)
	//// 1判断是否有该appName的更新记录 无则判断无需更新
	var data *entity.ClientUpdate
	r, err := dao.ClientUpdate.Ctx(ctx).Where(g.Map{
		"app_name": req.AppName,
	}).OrderDesc("created_at").One()
	if err != nil {
		return
	}
	err = r.Struct(&data)
	g.Log().Debug(ctx, data)
	if g.IsNil(data) {
		res.NeedUpdate = false

		return
	}
	g.Log().Debug(ctx, data.AppVersion, req.AppVersion)
	// 判断版本是否相同
	if data.AppVersion == req.AppVersion {
		res.NeedUpdate = false
		return
	}
	res.NeedUpdate = true
	res.Url = data.AppUrl
	res.AppVersion = data.AppVersion
	return
}

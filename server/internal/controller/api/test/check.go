package testcheck

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	testcheck "hotgo/api/api/test"
	"hotgo/internal/service"
)

var (
	TestCheck = cTestCheck{}
)

type cTestCheck struct {
}

func (c *cTestCheck) Add(ctx context.Context, req *testcheck.AddReq) (res *testcheck.AddRes, err error) {
	g.Log().Debug(ctx, "控制层", req)
	err = service.TestCheck().Add(ctx, &req.CheckAddInp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

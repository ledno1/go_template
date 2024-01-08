package base

import (
	"context"
	"hotgo/api/home/base"
	"hotgo/internal/service"
)

var Pay = cPay{}

type cPay struct{}

func (c *cPay) Index(ctx context.Context, _ *base.PayIndexReq) (res *base.PayIndexRes, err error) {
	service.View().RenderTpl(ctx, "pay/pay.html")
	return
}

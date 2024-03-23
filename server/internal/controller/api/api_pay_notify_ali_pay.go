package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/api/api/pay"
)

func (c *ControllerPay) NotifyAliPay(ctx context.Context, req *pay.NotifyAliPayReq) (res *pay.NotifyAliPayRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

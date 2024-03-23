package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/api/api/pay"
)

func (c *ControllerPay) NotifyWxPay(ctx context.Context, req *pay.NotifyWxPayReq) (res *pay.NotifyWxPayRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

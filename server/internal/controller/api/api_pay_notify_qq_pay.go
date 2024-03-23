package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/api/api/pay"
)

func (c *ControllerPay) NotifyQQPay(ctx context.Context, req *pay.NotifyQQPayReq) (res *pay.NotifyQQPayRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

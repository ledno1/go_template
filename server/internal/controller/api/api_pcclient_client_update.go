package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/api/api/pcclient"
)

func (c *ControllerPcclient) ClientUpdate(ctx context.Context, req *pcclient.ClientUpdateReq) (res *pcclient.ClientUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

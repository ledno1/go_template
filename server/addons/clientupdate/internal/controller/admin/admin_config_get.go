package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/config"
)

func (c *ControllerConfig) Get(ctx context.Context, req *config.GetReq) (res *config.GetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/config"
)

func (c *ControllerConfig) Update(ctx context.Context, req *config.UpdateReq) (res *config.UpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

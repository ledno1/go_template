package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/clientupdate"
)

func (c *ControllerClientupdate) View(ctx context.Context, req *clientupdate.ViewReq) (res *clientupdate.ViewRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

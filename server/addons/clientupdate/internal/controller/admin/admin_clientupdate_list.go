package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/clientupdate"
)

func (c *ControllerClientupdate) List(ctx context.Context, req *clientupdate.ListReq) (res *clientupdate.ListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

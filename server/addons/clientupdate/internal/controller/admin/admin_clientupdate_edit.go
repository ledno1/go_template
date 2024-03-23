package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/clientupdate"
)

func (c *ControllerClientupdate) Edit(ctx context.Context, req *clientupdate.EditReq) (res *clientupdate.EditRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

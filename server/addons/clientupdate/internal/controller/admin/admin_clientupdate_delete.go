package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/clientupdate"
)

func (c *ControllerClientupdate) Delete(ctx context.Context, req *clientupdate.DeleteReq) (res *clientupdate.DeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

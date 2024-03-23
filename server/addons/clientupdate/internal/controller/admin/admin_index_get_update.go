package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/admin/index"
)

func (c *ControllerIndex) GetUpdate(ctx context.Context, req *index.GetUpdateReq) (res *index.GetUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

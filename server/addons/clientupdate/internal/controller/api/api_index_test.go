package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/addons/clientupdate/api/api/index"
)

func (c *ControllerIndex) Test(ctx context.Context, req *index.TestReq) (res *index.TestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

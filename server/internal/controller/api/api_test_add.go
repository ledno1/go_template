package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/api/api/test"
)

func (c *ControllerTest) Add(ctx context.Context, req *test.AddReq) (res *test.AddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

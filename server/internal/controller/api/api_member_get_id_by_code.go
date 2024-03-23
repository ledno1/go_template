package api

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hotgo/api/api/member"
)

func (c *ControllerMember) GetIdByCode(ctx context.Context, req *member.GetIdByCodeReq) (res *member.GetIdByCodeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

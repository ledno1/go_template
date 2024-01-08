package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/testin"
	"hotgo/internal/service"
)

type sTestCheck struct{}

func NewTestCheck() *sTestCheck {
	return &sTestCheck{}
}

func init() {
	service.RegisterTestCheck(NewTestCheck())
}

func (s *sTestCheck) Add(ctx context.Context, in *testin.CheckAddInp) (err error) {
	g.Log().Debug(ctx, "服务逻辑层", in)

	return
}

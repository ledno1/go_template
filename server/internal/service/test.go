package service

import (
	"context"
	"hotgo/internal/model/input/testin"
)

type (
	ITestCheck interface {
		Add(ctx context.Context, in *testin.CheckAddInp) (err error)
	}
)

var (
	localTestCheck ITestCheck
)

func TestCheck() ITestCheck {
	if localTestCheck == nil {
		panic("implement not found for interface ISysProvinces, forgot register?")
	}
	return localTestCheck
}

func RegisterTestCheck(i ITestCheck) {
	localTestCheck = i
}

package testcheck

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/testin"
)

type AddReq struct {
	g.Meta `path:"/testcheck/add" mothod:"post" tags:"测试" summary:"测试"`
	testin.CheckAddInp
}

type AddRes struct {
}

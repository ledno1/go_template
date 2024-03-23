package pcclient

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/clientin"
)

// 客户端获取更新
type ClientUpdateReq struct {
	g.Meta `path:"/pcclient/clientupdate" method:"post" tags:"" summary:""`
	*clientin.ClientUpdateInp
}
type ClientUpdateRes struct {
	NeedUpdate bool   `json:"needUpdate"`
	Url        string `json:"url"`
	AppVersion string `json:"appVersion"`
}

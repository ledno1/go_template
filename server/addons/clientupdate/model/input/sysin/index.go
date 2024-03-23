// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// IndexTestInp 测试
type IndexTestInp struct {
	Name string `json:"name" d:"HotGo" dc:"名称"`
}

func (in *IndexTestInp) Filter(ctx context.Context) (err error) {
	return
}

type IndexTestModel struct {
	Name   string      `json:"name" dc:"名称"`
	Module string      `json:"module" dc:"当前插件模块"`
	Time   *gtime.Time `json:"time" dc:"当前时间"`
}

type GetUpdateInp struct {
	AppName    int    `json:"appName"    dc:"应用类型"`
	AppVersion string `json:"appVersion" dc:"应用版本"`
}

//	func (in *GetUpdateInp) Filter(ctx context.Context) (err error) {
//		if in.AppName == "" {
//			return gerror.New("应用名称不能为空")
//		}
//		// 长度
//		if len(in.AppName) > 32 {
//			return gerror.New("应用名称不能超过32个字符")
//		}
//		return
//	}
type GetUpdateModel struct {
	NeedUpdate bool   `json:"needUpdate" dc:"是否需要更新"`
	Url        string `json:"url" dc:"更新下载地址"`
}

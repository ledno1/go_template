// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package common

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/storager"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/format"
)

type sCommonUpload struct{}

func NewCommonUpload() *sCommonUpload {
	return &sCommonUpload{}
}

func init() {
	service.RegisterCommonUpload(NewCommonUpload())
}

// UploadFile 上传文件
func (s *sCommonUpload) UploadFile(ctx context.Context, uploadType string, file *ghttp.UploadFile) (res *sysin.AttachmentListModel, err error) {
	attachment, err := storager.DoUpload(ctx, uploadType, file)
	if err != nil {
		return
	}
	filePath := attachment.FileUrl

	attachment.FileUrl = storager.LastUrl(ctx, attachment.FileUrl, attachment.Drive)
	res = &sysin.AttachmentListModel{
		SysAttachment: *attachment,
		SizeFormat:    format.FileSize(attachment.Size),
	}
	g.Log().Debug(ctx, file.Filename)
	// 分割file.Filename 获取版本号
	fileName := gstr.Split(file.Filename, ".")

	fileHax := gstr.Split(gconv.String(fileName[0]), "-")
	// 新增版本号
	var (
		uf entity.ClientUpdate
		cf *entity.SysConfig
	)
	uf.AppName = 1
	//// 从配置服务器中获取服务器域名
	_ = dao.SysConfig.Ctx(ctx).Where(g.Map{"name": "网站域名"}).Data().Scan(&cf)
	g.Log().Debug(ctx, cf)
	uf.AppUrl = cf.Value + "/" + filePath
	g.Log().Debug(ctx, fileHax)
	g.Log().Debug(ctx, fileHax[1])
	uf.AppVersion = fileHax[1]
	// 插入数据
	_, _ = dao.ClientUpdate.Ctx(ctx).Data(uf).Insert()
	return
}

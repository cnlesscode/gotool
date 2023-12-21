package gintool

import (
	"errors"
	"strings"
	"time"

	"github.com/cnlesscode/gotool/gfs"
	"github.com/cnlesscode/gotool/gmd5"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Upload struct {
	FileName         string
	MaxSize          int64
	AllowExeNames    string
	AllowTypes       string
	TargetDir        string
	DirNamingRule    string
	FileNamingRule   string
	UploadedFilePath string
}

func (m *Upload) Run(ctx *gin.Context) (error, string) {
	// 检查文件名及文件有效性
	if m.FileName == "" {
		return errors.New("请设置上传文件名"), ""
	}
	file, err := ctx.FormFile(m.FileName)
	if err != nil {
		return errors.New("上传文件数据错误"), ""
	}
	// 检查文件大小
	if m.MaxSize < 1 {
		m.MaxSize = 10
	}
	if file.Size > m.MaxSize*1048576 {
		return errors.New("上传文件过大"), ""
	}
	// 检查文件类型
	if m.AllowTypes == "" {
		return errors.New("请设置允许的上传类型"), ""
	}
	fileTypes := file.Header["Content-Type"]
	if len(fileTypes) < 1 {
		return errors.New("上传文件类型识别错误"), ""
	}
	fileType := fileTypes[0]
	if m.AllowTypes != "" && m.AllowTypes != "*" {
		if !strings.Contains(m.AllowTypes, fileType) {
			return errors.New("上传文件类型错误"), ""
		}
	}
	// 检查扩展名
	extendName := strings.ToLower(gfs.GetExtension(file.Filename))
	if m.AllowExeNames != "" && m.AllowExeNames != "*" {
		if !strings.Contains(m.AllowExeNames, extendName) {
			return errors.New("上传文件扩展名错误"), ""
		}
	}
	// 检查文件夹是否存在，不存在则创建
	if m.TargetDir == "" {
		return errors.New("请设置上传目录"), ""
	}
	// 根据文件夹命名规则命名文件夹
	uploadTargetDir := ""
	if m.DirNamingRule == "" {
		uploadTargetDir = m.TargetDir + "/"
	} else if m.DirNamingRule == "year" {
		uploadTargetDir = m.TargetDir + time.Now().Format("2006") + "/"
	} else if m.DirNamingRule == "month" {
		uploadTargetDir = m.TargetDir + time.Now().Format("2006-01") + "/"
	} else if m.DirNamingRule == "day" {
		uploadTargetDir = m.TargetDir + time.Now().Format("2006-01-02") + "/"
	} else {
		uploadTargetDir = m.TargetDir + time.Now().Format("2006-01") + "/"
	}
	if !gfs.DirExists(uploadTargetDir) {
		err = gfs.MakeDir(uploadTargetDir)
		if err != nil {
			return errors.New("目标文件夹创建失败"), ""
		}
	}
	// 根据文件命名规则命名文件
	uploadTargetFile := ""
	if m.FileNamingRule == "" {
		uploadTargetFile = file.Filename
	} else if m.FileNamingRule == "random" {
		uploadTargetFile = gmd5.Md5(uuid.New().String()) + "." + extendName
	}
	// 上传文件至指定的完整文件路径
	err = ctx.SaveUploadedFile(file, uploadTargetDir+uploadTargetFile)
	if err == nil {
		return nil, uploadTargetDir + uploadTargetFile
	} else {
		return err, ""
	}
}

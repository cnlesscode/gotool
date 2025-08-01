package gintool

import (
	"errors"
	"mime/multipart"
	"os"
	"path"
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
	ExtendName       string
	AllowTypes       string
	FileType         string
	TargetDir        string
	DirNamingRule    string
	FileNamingRule   string
	UploadedFilePath string
	SourceFile       *multipart.FileHeader
}

func (m *Upload) Run(ctx *gin.Context) (string, error) {
	var err error

	// 检查文件名及文件有效性
	if m.FileName == "" {
		return "", errors.New("请设置上传文件名")
	}
	m.SourceFile, err = ctx.FormFile(m.FileName)
	if err != nil {
		return "", errors.New("上传文件数据错误")
	}
	// 检查文件大小
	if m.MaxSize < 1 {
		m.MaxSize = 10
	}
	if m.SourceFile.Size > m.MaxSize*1048576 {
		return "", errors.New("上传文件过大")
	}

	// 检查文件类型
	// * 代表全部
	if m.AllowTypes != "" && m.AllowTypes != "*" {
		fileTypes := m.SourceFile.Header["Content-Type"]
		if len(fileTypes) < 1 {
			return "", errors.New("上传文件类型识别错误")
		}
		m.FileType = fileTypes[0]
		if !MatchMimeType(m.AllowTypes, m.FileType) {
			return "", errors.New("上传文件类型错误")
		}
	}

	// 检查扩展名
	// * 代表全部
	m.ExtendName = strings.ToLower(path.Ext(m.SourceFile.Filename))
	m.ExtendName = strings.TrimPrefix(m.ExtendName, ".")
	if m.AllowExeNames != "" && m.AllowExeNames != "*" {
		if !strings.Contains(m.AllowExeNames, m.ExtendName) {
			return "", errors.New("上传文件扩展名错误")
		}
	}

	if m.TargetDir == "" {
		return "", errors.New("请设置上传目录")
	}

	// 上传目录应以 ./开头
	if !strings.HasPrefix(m.TargetDir, "./") {
		return "", errors.New("上传目录格式错误，应以 ./ 开头")
	}

	// 根据文件夹命名规则命名文件夹
	uploadTargetDir := ""
	switch m.DirNamingRule {
	case "":
		uploadTargetDir = m.TargetDir
	case "year":
		uploadTargetDir = path.Join(m.TargetDir, time.Now().Format("2006"))
	case "month":
		uploadTargetDir = path.Join(m.TargetDir, time.Now().Format("200601"))
	case "day":
		uploadTargetDir = path.Join(m.TargetDir, time.Now().Format("20060102"))
	default:
		uploadTargetDir = path.Join(m.TargetDir, time.Now().Format("200601"))
	}
	// 检查文件夹是否存在，不存在则创建
	if !gfs.DirExists(uploadTargetDir) {
		err = os.MkdirAll(uploadTargetDir, 0777)
		if err != nil {
			return "", errors.New("目标文件夹创建失败")
		}
	}

	// 根据文件命名规则命名文件
	uploadTargetFile := ""
	switch m.FileNamingRule {
	case "":
		uploadTargetFile = m.SourceFile.Filename
	case "random":
		uploadTargetFile = gmd5.Md5(uuid.New().String()) + "." + m.ExtendName
	}

	// 上传文件至指定的完整文件路径
	m.UploadedFilePath = path.Join(uploadTargetDir, uploadTargetFile)
	err = ctx.SaveUploadedFile(m.SourceFile, m.UploadedFilePath)
	if err == nil {
		return m.UploadedFilePath, nil
	} else {
		return "", err
	}
}

// MimeType 检查
func MatchMimeType(allowedTypes, fileType string) bool {

	// 精确匹配
	if allowedTypes == fileType {
		return true
	}

	// 检查是否使用了通配符格式（如 image/*）
	if strings.HasSuffix(allowedTypes, "/*") {
		prefix := strings.TrimSuffix(allowedTypes, "/*")
		return strings.HasPrefix(fileType, prefix+"/")
	}

	// 检查是否在逗号分隔的列表中
	types := strings.Split(allowedTypes, ",")
	for _, t := range types {
		t = strings.TrimSpace(t)
		if t == fileType {
			return true
		}
		// 支持通配符检查
		if strings.HasSuffix(t, "/*") {
			prefix := strings.TrimSuffix(t, "/*")
			if strings.HasPrefix(fileType, prefix+"/") {
				return true
			}
		}
	}
	return false
}

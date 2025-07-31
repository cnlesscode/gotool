package static

import (
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Endpoint     string
	KeyId        string
	Secret       string
	BucketName   string
	LocalBaseUrl string
	BaseUrl      string
	Type         string
}

type StaticCloud interface {
	InitClient(config Config)
	UploadFile(fileUrl string) error
	DownloadFile(fileUrl string) error
	RemoveFile(fileUrl string, removeLocalFile bool) error
	MoveFile(fileUrl string, targerDir string) (string, error)
}

// 规划文件路径
func InitFileUrl(fileUrl string, config Config) (localUrl string, cloudUrl string) {
	if fileUrl[0:2] == "./" {
		localUrl = fileUrl
		cloudUrl = fileUrl[2:]
	} else if fileUrl[0:1] == "/" {
		localUrl = "." + fileUrl
		cloudUrl = fileUrl[1:]
	} else if fileUrl[0:4] == "http" {
		cloudUrl = strings.ReplaceAll(fileUrl, config.BaseUrl, "")
		localUrl = "./" + cloudUrl
	} else {
		localUrl = fileUrl
		cloudUrl = fileUrl
	}
	return localUrl, cloudUrl
}

func New(config Config) StaticCloud {
	var staticCloud StaticCloud
	switch config.Type {
	case "AliOSS":
		staticCloud = &AliOSS{}
	case "TencentCOS":
		staticCloud = &TencentCOS{}
	default:
		staticCloud = &Local{}
	}
	staticCloud.InitClient(config)
	return staticCloud
}

func MoveFile(fileUrl string, targerUrl string) error {
	// 检查目标目录是否存在，不存在则创建
	targerDir := filepath.Dir(targerUrl)
	if _, err := os.Stat(targerDir); os.IsNotExist(err) {
		err := os.MkdirAll(targerDir, 0777)
		if err != nil {
			return err
		}
	}
	// 移动文件至目标文件夹
	return os.Rename(fileUrl, targerUrl)
}

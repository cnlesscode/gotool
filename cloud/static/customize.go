package static

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/cnlesscode/gotool/ip"
	"github.com/cnlesscode/gotool/request"
)

type Customize struct {
	Config       Config
	IsSelfServer bool
}

func (m *Customize) InitClient(config Config) {
	m.Config = config
	u, err := url.Parse(m.Config.Endpoint)
	if err != nil {
		return
	}
	host := u.Hostname()
	if host == "localhost" || host == "127.0.0.1" || host == "::1" {
		m.IsSelfServer = true
		return
	}

	// 获取本机 ip
	ip := ip.GetLocalIP()

	if ip == host {
		m.IsSelfServer = true
	}

}

// Remove File
func (m *Customize) RemoveFile(fileUrl string, removeLocalFile bool) error {
	fileUrl, _ = InitFileUrl(fileUrl, m.Config)
	if removeLocalFile {
		os.Remove(fileUrl)
	}
	if m.IsSelfServer {
		return nil
	}
	response, err := request.POST(
		m.Config.Endpoint,
		map[string]string{
			"Secret":   m.Config.Secret,
			"FilePath": fileUrl,
			"Action":   "Remove",
		},
		nil,
	)
	if err != nil {
		return err
	}

	var res map[string]string = map[string]string{}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return err
	}
	if res["ErrorCode"] == "0" {
		return nil
	}
	return errors.New(res["Data"])
}

func (m *Customize) UploadFile(fileUrl string) error {
	if m.IsSelfServer {
		return nil
	}
	fileUrl, _ = InitFileUrl(fileUrl, m.Config)
	// 读取本地文件，将文件以post 数据形式发给云服务器
	file, err := os.Open(fileUrl)
	if err != nil {
		return fmt.Errorf("open file %s: %w", fileUrl, err)
	}
	defer file.Close()

	return request.UploadFile(
		m.Config.Endpoint,
		fileUrl, map[string]string{
			"Secret":   m.Config.Secret,
			"FilePath": fileUrl,
			"Action":   "UploadFile",
		})
}

func (m *Customize) DownloadFile(fileUrl string) error {
	return nil
}

func (m *Customize) MoveFile(fileUrl string, targetUrl string) (string, error) {
	// 本地移动
	fileUrl, _ = InitFileUrl(fileUrl, m.Config)
	targetUrl, _ = InitFileUrl(targetUrl, m.Config)
	err := MoveFile(fileUrl, targetUrl)
	if err != nil {
		return "", err
	}

	if m.IsSelfServer {
		return targetUrl, nil
	}

	// 云移动
	fileUrl, _ = InitFileUrl(fileUrl, m.Config)
	targetUrl, _ = InitFileUrl(targetUrl, m.Config)

	// 发送移动请求
	response, err := request.POST(
		m.Config.Endpoint,
		map[string]string{
			"Secret":        m.Config.Secret,
			"FilePath":      fileUrl,
			"TargetFileUrl": targetUrl,
			"Action":        "MoveFile",
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	var res map[string]string = map[string]string{}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return "", err
	}
	if res["ErrorCode"] == "0" {
		return targetUrl, nil
	}
	return "", errors.New(res["Data"])
}

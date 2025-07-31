package static

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentCOS struct {
	Client *cos.Client
	Ready  bool
	Config Config
}

// 初始化客户端
func (m *TencentCOS) InitClient(config Config) {
	m.Ready = false
	m.Config = config
	u, _ := url.Parse(m.Config.BucketName)
	b := &cos.BaseURL{BucketURL: u}
	m.Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。
			// 子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretID: m.Config.KeyId,
			// 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。
			// 子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: m.Config.Secret,
		},
	})
	m.Ready = true
}

// 上传文件
func (m *TencentCOS) UploadFile(fileUrl string) error {
	if !m.Ready {
		return errors.New("not ready")
	}
	localFileUrl, cloudFileUrl := InitFileUrl(fileUrl, m.Config)
	_, err := m.Client.Object.PutFromFile(context.Background(), cloudFileUrl, localFileUrl, nil)
	return err
}

// 删除文件
func (m *TencentCOS) RemoveFile(fileUrl string, removeLocalFile bool) error {
	if !m.Ready {
		return errors.New("not ready")
	}
	localFileUrl, cloudFileUrl := InitFileUrl(fileUrl, m.Config)
	if removeLocalFile {
		os.Remove(localFileUrl)
	}
	_, err := m.Client.Object.Delete(context.Background(), cloudFileUrl)
	return err
}

// 下载文件
func (m *TencentCOS) DownloadFile(fileUrl string) error {
	if !m.Ready {
		return errors.New("not ready")
	}
	localFileUrl, cloudFileUrl := InitFileUrl(fileUrl, m.Config)
	_, err := m.Client.Object.GetToFile(context.Background(), cloudFileUrl, localFileUrl, nil)
	return err
}

func (m *TencentCOS) MoveFile(fileUrl string, targerDir string) (string, error) {

	// 本地移动
	fileUrl, _ = InitFileUrl(fileUrl, m.Config)
	targerUrl := path.Join(targerDir, filepath.Base(fileUrl))
	targerUrl, _ = InitFileUrl(targerUrl, m.Config)
	MoveFile(fileUrl, targerUrl)

	// 腾讯云移动
	_, fileUrl = InitFileUrl(fileUrl, m.Config)
	_, targerUrl = InitFileUrl(targerUrl, m.Config)

	// 移动对象
	var err error
	bucketURL := m.Client.BaseURL.BucketURL
	sourceURL := fmt.Sprintf("%s/%s", bucketURL.Host, fileUrl)
	_, _, err = m.Client.Object.Copy(context.Background(), targerUrl, sourceURL, nil)
	if err != nil {
		return "", err
	}
	_, err = m.Client.Object.Delete(context.Background(), fileUrl, nil)
	if err != nil {
		return "", err
	}
	return targerUrl, nil
}

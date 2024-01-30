package static

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentCOS struct {
	BucketURL string
	SecretId  string
	SecretKey string
	BaseUrl   string
}

// 初始化客户端
func (m TencentCOS) InitClient() *cos.Client {
	u, _ := url.Parse(m.BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。
			// 子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretID: m.SecretId,
			// 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。
			// 子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: m.SecretKey,
		},
	})
}

// 上传文件
func (m *TencentCOS) UploadFile(cloudFileUrl string, localFileUrl string) error {
	c := m.InitClient()
	_, err := c.Object.PutFromFile(context.Background(), cloudFileUrl, localFileUrl, nil)
	return err
}

// 删除文件
func (m *TencentCOS) RemoveFile(cloudFileUrl string, localFileUrl string, removeLocalFile bool) error {
	if removeLocalFile {
		os.Remove(localFileUrl)
	}
	c := m.InitClient()
	_, err := c.Object.Delete(context.Background(), cloudFileUrl)
	return err
}

package static

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 阿里云 OSS 文档
// https://help.aliyun.com/zh/oss/developer-reference/introduction-3?spm=5176.8466032.console-base_help.dexternal.67181450WWM1Nx

type AliOSS struct {
	Client *oss.Client
	Ready  bool
	Config Config
	Bucket *oss.Bucket
}

func (m *AliOSS) InitClient(config Config) {
	m.Ready = false
	m.Config = config
	client, err := oss.New(config.Endpoint, config.KeyId, config.Secret)
	if err != nil {
		return
	}
	m.Client = client
	m.Bucket, err = client.Bucket(config.BucketName)
	if err != nil {
		return
	}
	m.Ready = true
}

// Remove File
func (m *AliOSS) RemoveFile(fileUrl string, removeLocalFile bool) error {
	fileUrl, cloudFileUrl := InitFileUrl(fileUrl, m.Config)
	if !m.Ready {
		return errors.New("not ready")
	}
	if removeLocalFile {
		os.Remove(fileUrl)
	}
	return m.Bucket.DeleteObject(cloudFileUrl)
}

// Upload File
func (m *AliOSS) UploadFile(fileUrl string) error {
	if !m.Ready {
		return errors.New("not ready")
	}
	fileUrl, cloudFileUrl := InitFileUrl(fileUrl, m.Config)
	return m.Bucket.PutObjectFromFile(cloudFileUrl, fileUrl)
}

// 下载文件
func (m *AliOSS) DownloadFile(fileUrl string) error {
	fileUrl, cloudFileUrl := InitFileUrl(fileUrl, m.Config)
	if !m.Ready {
		return errors.New("not ready")
	}
	// 检查对象是否存在
	_, err := m.Bucket.GetObjectMeta(cloudFileUrl)
	if err != nil {
		return err
	}
	// 下载文件
	return m.Bucket.GetObjectToFile(cloudFileUrl, fileUrl)
}

func (m *AliOSS) MoveFile(fileUrl string, targerDir string) (string, error) {

	// 本地移动
	fileUrl, _ = InitFileUrl(fileUrl, m.Config)
	targerUrl := path.Join(targerDir, filepath.Base(fileUrl))
	targerUrl, _ = InitFileUrl(targerUrl, m.Config)
	MoveFile(fileUrl, targerUrl)

	// 阿里云移动
	// 01. 拷贝
	_, fileUrl = InitFileUrl(fileUrl, m.Config)
	_, targerUrl = InitFileUrl(targerUrl, m.Config)
	_, err := m.Bucket.CopyObject(fileUrl, targerUrl)
	if err != nil {
		return "", err
	}
	// 02. 删除
	err = m.Bucket.DeleteObject(fileUrl)
	if err != nil {
		return "", err
	}
	return targerUrl, nil
}

package static

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliOSS struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	BaseUrl         string
}

func (aliOSS AliOSS) InitClient() (*oss.Client, error) {
	client, err := oss.New(aliOSS.Endpoint, aliOSS.AccessKeyId, aliOSS.AccessKeySecret)
	return client, err
}

// Remove AliOSS File
func (aliOSS AliOSS) RemoveFile(cloudFileUrl string, localFileUrl string, removeLocalFile bool) error {
	client, err := aliOSS.InitClient()
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(aliOSS.BucketName)
	if err != nil {
		return err
	}
	if removeLocalFile {
		os.Remove(localFileUrl)
	}
	return bucket.DeleteObject(cloudFileUrl)
}

// upload a file to aliOSS
func (aliOSS AliOSS) UploadFile(cloudFileUrl string, localFileUrl string) error {
	client, err := aliOSS.InitClient()
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(aliOSS.BucketName)
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(cloudFileUrl, localFileUrl)
}

// 下载文件
func (aliOSS AliOSS) DownloadFile(cloudFileUrl string, localFileUrl string) error {
	client, err := aliOSS.InitClient()
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(aliOSS.BucketName)
	if err != nil {
		return err
	}
	return bucket.GetObjectToFile(cloudFileUrl, localFileUrl)
}

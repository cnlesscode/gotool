package cloud

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

func (aliOSS *AliOSS) InitClient() (*oss.Client, error) {
	client, err := oss.New(aliOSS.Endpoint, aliOSS.AccessKeyId, aliOSS.AccessKeySecret)
	return client, err
}

// Remove AliOSS File
func (aliOSS *AliOSS) RemoveFile(fileUrl string, removeLocalFile bool) error {
	if removeLocalFile {
		os.Remove(fileUrl)
	}
	if aliOSS.BaseUrl != "/" {
		client, err := aliOSS.InitClient()
		if err != nil {
			return err
		}
		bucket, err := client.Bucket(aliOSS.BucketName)
		if err != nil {
			return err
		}
		fileUrlAliOss := ""
		if fileUrl[0:2] == "./" {
			fileUrlAliOss = fileUrl[2:]
		} else {
			fileUrlAliOss = fileUrl
		}
		err = bucket.DeleteObject(fileUrlAliOss)
		if err != nil {
			return err
		}
	}
	return nil
}

// upload a file to aliOSS
func (aliOSS *AliOSS) UploadFile(fileUrl string) error {
	if aliOSS.BaseUrl == "/" {
		return nil
	}
	client, err := aliOSS.InitClient()
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(aliOSS.BucketName)
	if err != nil {
		return err
	}
	fileUrlAliOss := ""
	if fileUrl[0:2] == "./" {
		fileUrlAliOss = fileUrl[2:]
	} else {
		fileUrlAliOss = fileUrl
	}
	err = bucket.PutObjectFromFile(fileUrlAliOss, fileUrl)
	if err != nil {
		return err
	}
	return nil
}

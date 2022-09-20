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

// Remove AliOSS File
func (aliOSS *AliOSS) RemoveOSSFile(fileUrl string, removeLocalFile bool) error {

	if removeLocalFile {
		os.Remove(fileUrl)
	}

	if aliOSS.BaseUrl != "/" {
		client, err := oss.New(aliOSS.Endpoint, aliOSS.AccessKeyId, aliOSS.AccessKeySecret)
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
func (aliOSS *AliOSS) FileToOSS(fileUrl string) error {

	if aliOSS.BaseUrl == "/" {
		return nil
	}

	client, err := oss.New(aliOSS.Endpoint, aliOSS.AccessKeyId, aliOSS.AccessKeySecret)
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

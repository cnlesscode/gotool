package static

import "strings"

type Config struct {
	Endpoint   string
	KeyId      string
	Secret     string
	BucketName string
	BaseUrl    string
	Type       string
}

// 上传文件到云存储
func UploadFile(config *Config, fileUrl string) error {
	localFileUrl, cloudFileUrl := InitFileUrl(fileUrl, config)
	if config.Type == "Local" {
		st := Local{
			BaseUrl: config.BaseUrl,
		}
		return st.UploadFile(cloudFileUrl)
	} else if config.Type == "AliOSS" {
		st := AliOSS{
			Endpoint:        config.Endpoint,
			AccessKeyId:     config.KeyId,
			AccessKeySecret: config.Secret,
			BucketName:      config.BucketName,
			BaseUrl:         config.BaseUrl,
		}
		return st.UploadFile(cloudFileUrl, localFileUrl)
	} else if config.Type == "TencentCOS" {
		st := TencentCOS{
			BucketURL: config.Endpoint,
			SecretId:  config.KeyId,
			SecretKey: config.Secret,
			BaseUrl:   config.BaseUrl,
		}
		return st.UploadFile(cloudFileUrl, localFileUrl)
	}
	return nil
}

// 上传文件到云存储
func RemoveFile(config *Config, fileUrl string, removeLocalFile bool) error {
	localFileUrl, cloudFileUrl := InitFileUrl(fileUrl, config)
	if config.Type == "Local" {
		st := Local{
			BaseUrl: config.BaseUrl,
		}
		return st.RemoveFile(localFileUrl)
	} else if config.Type == "AliOSS" {
		st := AliOSS{
			Endpoint:        config.Endpoint,
			AccessKeyId:     config.KeyId,
			AccessKeySecret: config.Secret,
			BucketName:      config.BucketName,
			BaseUrl:         config.BaseUrl,
		}
		return st.RemoveFile(cloudFileUrl, localFileUrl, true)
	} else if config.Type == "TencentCOS" {
		st := TencentCOS{
			BucketURL: config.Endpoint,
			SecretId:  config.KeyId,
			SecretKey: config.Secret,
			BaseUrl:   config.BaseUrl,
		}
		return st.RemoveFile(cloudFileUrl, localFileUrl, true)
	}
	return nil
}

// 规划图片文件路径
func InitFileUrl(fileUrl string, config *Config) (string, string) {
	localUrl := ""
	cloudUrl := ""
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

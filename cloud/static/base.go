package static

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
	if config.Type == "Local" {
		st := Local{
			BaseUrl: config.BaseUrl,
		}
		return st.UploadFile(fileUrl)
	} else if config.Type == "AliOSS" {
		st := AliOSS{
			Endpoint:        config.Endpoint,
			AccessKeyId:     config.KeyId,
			AccessKeySecret: config.Secret,
			BucketName:      config.BucketName,
			BaseUrl:         config.BaseUrl,
		}
		return st.UploadFile(fileUrl)
	} else if config.Type == "TencentCOS" {
		st := TencentCOS{
			BucketURL: config.Endpoint,
			SecretId:  config.KeyId,
			SecretKey: config.Secret,
			BaseUrl:   config.BaseUrl,
		}
		return st.UploadFile(fileUrl)
	}
	return nil
}

// 上传文件到云存储
func RemoveFile(config *Config, fileUrl string, removeLocalFile bool) error {
	if config.Type == "Local" {
		st := Local{
			BaseUrl: config.BaseUrl,
		}
		return st.RemoveFile(fileUrl)
	} else if config.Type == "AliOSS" {
		st := AliOSS{
			Endpoint:        config.Endpoint,
			AccessKeyId:     config.KeyId,
			AccessKeySecret: config.Secret,
			BucketName:      config.BucketName,
			BaseUrl:         config.BaseUrl,
		}
		return st.RemoveFile(fileUrl, removeLocalFile)
	} else if config.Type == "TencentCOS" {
		st := TencentCOS{
			BucketURL: config.Endpoint,
			SecretId:  config.KeyId,
			SecretKey: config.Secret,
			BaseUrl:   config.BaseUrl,
		}
		return st.RemoveFile(fileUrl, removeLocalFile)
	}
	return nil
}

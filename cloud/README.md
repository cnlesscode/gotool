### 云操作工具包

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/cloud"
)
```

### 阿里云静态云存储

#### FileToOSS(fileUrl string) error
函数功能 : 将本地文件同步到阿里云 OSS
返回格式 : error
```go
func main() {
	aliOSS := cloud.AliOSS{
		// 云存储节点
		Endpoint: "oss-cn-beijing.aliyuncs.com",
		// AccessKeyId
		AccessKeyId: "************",
		// AccessKeySecret
		AccessKeySecret: "************",
		// 空间名称
		BucketName: "staticfilesali",
		// 以 / 结尾  / 代表关闭云存储
		BaseUrl: "https://静态云存储基础路径/",
	}
	err := aliOSS.FileToOSS("./1.txt")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
}
```

#### RemoveOSSFile(fileUrl string, removeLocalFile bool) error
函数功能 : 删除 OSS 文件及本地文件
返回格式 : error
```go
func main() {
	aliOSS := cloud.AliOSS{
		// 云存储节点
		Endpoint: "oss-cn-beijing.aliyuncs.com",
		// AccessKeyId
		AccessKeyId: "************",
		// AccessKeySecret
		AccessKeySecret: "************",
		// 空间名称
		BucketName: "staticfilesali",
		// 以 / 结尾  / 代表关闭云存储
		BaseUrl: "https://静态云存储基础路径/",
	}
	err := aliOSS.RemoveOSSFile("./1.txt", true)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
}
```


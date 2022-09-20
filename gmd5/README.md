### MD5 工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/gmd5"
)
```

#### Md5(原始字符 string) string
函数功能 : 简单 md5 加密
返回格式 : string
```go
fmt.Printf("gmd5.Md5(\"test\"): %v\n", gmd5.Md5("test"))
```

#### ToMd5Pwd(原始字符 string) string 
函数功能 : 复杂 md5 加密 (2次 Md5 加密后再进行混淆)
返回格式 : string
```go
fmt.Printf("gmd5.Md5(\"test\"): %v\n", gmd5.Md5("test"))
```

#### ToMd5Pwd(原始字符 string) string 
函数功能 : 复杂 md5 加密 (2次 Md5 加密后再进行混淆)
返回格式 : string
```go
fmt.Printf("gmd5.ToMd5Pwd(\"test\"): %v\n", gmd5.ToMd5Pwd("test"))
```

#### PwdToMd5(复杂Md5密码 string) string 
函数功能 : 将复杂 md5 密码翻译为 2次md5
返回格式 : string
```go
fmt.Printf("%v\n", gmd5.PwdToMd5("fb469d7ef430b0baf0cab6c436e701337529"))
```
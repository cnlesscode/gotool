### 文件操作工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/gfs"
)
```

#### PathExists(文件或文件夹路径 string) bool
函数功能 : 判断文件或者文件夹是否存在
返回格式 : bool
```go
fmt.Printf("%v\n", gfs.PathExists("./a.txt")) // true || false
```

#### FileExists(文件路径 string) bool
函数功能 : 判断文件是否存在
返回格式 : bool
```go
fmt.Printf("%v\n", gfs.FileExists("./a.txt")) // true || false
```

#### DirExists(文件夹路径 string) bool 
函数功能 : 判断文件夹是否存在
返回格式 : bool
```go
fmt.Printf("%v\n", gfs.DirExists("./fs")) // true || false
```

#### CopyFile(源文件路径,目标文件路径) error
函数功能 : 拷贝文件
返回格式 : error
```go
err := gfs.CopyFile("./a.txt", "./b.txt") // true
if err != nil {
	fmt.Printf("error: %v\n", err.Error())
}
```

#### CopyDir(源文件夹路径,目标文件夹路径) error
函数功能 : 拷贝文件夹
返回格式 : error
```go
err := gfs.CopyDir("./fs", "./a") // true
if err != nil {
	fmt.Printf("error: %v\n", err.Error())
}
```

#### MakeDir(文件夹路径 string) error
函数功能 : 创建文件夹
返回格式 : error
```go
err := gfs.MakeDir("./a/b/c") // true
if err != nil {
	fmt.Printf("error: %v\n", err.Error())
}
```

#### RemoveDir(文件夹路径 string) error
函数功能 : 删除文件夹
返回格式 : error
```go
err := gfs.RemoveDir("./a") // true
if err != nil {
	fmt.Printf("error: %v\n", err.Error())
}
```

#### DirSize(文件夹路径 string) (size, error)
函数功能 : 获取文件夹大小, 单位字节
返回格式 : 文件夹大小, error
```go
size, err := gfs.DirSize("./fs") // true
if err != nil {
	fmt.Printf("error: %v\n", err.Error())
} else {
	fmt.Printf("size: %v\n", size)
}
```

#### ModifyTime(文件/夹路径 string) (int64, error)
函数功能 : 获取文件夹或文件夹修改时间
返回格式 : 文件夹或文件夹修改时间, error
```go
modTime, err := gfs.ModifyTime("./fs")
if err != nil {
	fmt.Printf("err.Error(): %v\n", err.Error())
} else {
	fmt.Printf("modTime: %v\n", modTime)
}
```

#### ReadFile(文件路径 string) (string, error)
函数功能 : 读取文件内容
返回格式 : 文件内容, error
```go
content, err := gfs.ReadFile("./a.txt")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	} else {
		fmt.Printf("content: %v\n", content)
	}
}
```

#### WriteContentToFile(内容 string, 文件夹路径[*必须以/结尾] string, 文件名称 string) error
函数功能 : 向文件内写入数据 [ 自动创建文件 ]
返回格式 : error
```go
err := gfs.WriteContentToFile("content ...", "./a/b/c/", "ab.txt")
if err != nil {
	fmt.Printf("err.Error(): %v\n", err.Error())
}
```

#### AppendContentToFile(内容 string, 文件路径 string) error
函数功能 : 向文件尾部追加内容
返回格式 : error
```go
err := gfs.AppendContentToFile("\nappend content ...", "./ab.txt")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
}
```

#### PrependContentToFile(内容 string, 文件路径 string) error
函数功能 : 向文件头部部追加内容
返回格式 : error
```go
err := gfs.PrependContentToFile("Prepend content ...\n", "./ab.txt")
if err != nil {
	fmt.Printf("err.Error(): %v\n", err.Error())
}
```

#### GetExtension(文件路径 string) string
函数功能 : 获取文件扩展名
返回格式 : string
```go
fmt.Printf("%v\n", gfs.GetExtension("./a/b/c.tXt"))
```

#### GetFileName(文件路径 string, 分隔符) string
函数功能 : 从路径中获取文件名
返回格式 : string
```go
fileName := gfs.GetFileName("../a/b/c/rd.md", "/")
fmt.Printf("fileName: %v\n", fileName)
// fileName: rd.md
```
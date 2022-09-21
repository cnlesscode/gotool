### 字符串操作工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/string"
)
```

#### FirstUpper(s string) string
函数功能 : 首字母大写
返回格式 : string
```go
fmt.Printf("%v\n", gstring.FirstUpper("golang"))
```

#### FirstLower(s string) string
函数功能 : 首字母小写
返回格式 : string
```go
fmt.Printf("%v\n", gstring.FirstLower("Golang"))
```

#### FindImagesFromHtml(html string) []string
函数功能 : 从 html 中匹配出图片
返回格式 : []string
```go
html := `<div>...</div><img src="http://localhost/a.png" />`
images := gstring.FindImagesFromHtml(html)
fmt.Printf("images: %v\n", images)
```

#### TrimHtmlTags(html string) string
函数功能 : 去除 html 标签
返回格式 : string
```go
html := `<div>...</div><img src="http://localhost/a.png" />`
html = gstring.TrimHtmlTags(html)
fmt.Printf("%v\n", html)
```

### AnyToString
函数功能 : 任意类型转字符串
返回格式 : string
使用演示 :
```go
package main
import (
	"fmt"
	"github.com/cnlesscode/gotool/gstring"
)
func main() {
	a := 11
	aString := gstring.AnyToString(a)
	fmt.Printf("aString: %v\n", aString)
}
```

### StripPunctuation
函数功能 : 去除标点符号
参数说明 : 一段字符串
返回格式 : string
使用演示 :
```go
package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/gstring"
)

func main() {
	str := "hi! 您好啊~ 。"
	str = gstring.StripPunctuation(str)
	fmt.Printf("str: %v\n", str)
}
```

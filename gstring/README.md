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

### NLP 分词
```基于 github.com/go-ego/gse 的中文分词工具
字典位置 : "/resources/dict/zh/s_1.txt"
字典路径 : https://github.com/cnlesscode/gotool/tree/main/resources/dict
字典部署 : 请下载字典文件并部署到您的项目根目录 /resources 文件夹下
排除分词 : 您可以通过修改 /resources/dict/nplNotWords.txt 来排除不需要的分词
```
使用演示 :
```go
package main
import (
	"fmt"
	"github.com/cnlesscode/gotool/gstring"
)
func main() {
	NLPObject := gstring.NLP{Content: "测试分词文本 : 测试分词频率"}
	words := NLPObject.Cut()
	fmt.Printf("words: %v\n", words)
}
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

### nlp 分词工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/nlp"
)
```

#### 使用演示
```go
package main
import (
	"fmt"
	"github.com/cnlesscode/gotool/nlp"
)

func main() {
	words := nlp.Cut("我爱我的祖国")
	fmt.Printf("words: %v\n", words)
	words2 := nlp.Cut("坚持努力坚持执着")
	fmt.Printf("words: %v\n", words2)
}
```
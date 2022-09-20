### 数据分页工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/paginator"
)
```

#### 相关说明
```
1. 核心方法 paginator.Run
2. 返回值格式
type Pager struct {
	// 分页列表
	Pages       []int
	// 分页总数
	TotalPages  int
	// 首页
	FirstPage   int
	// 尾页
	PrePage     int
	// 下一页
	NextPage    int
	// 最后一页
	LastPage    int
	// 当前页
	CurrentPage int
}
3. 不需要展示的数据 对应值为 -1，前端可以利用此条件决定是否展示
```

#### 分页示例
```go
package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/paginator"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	pager := paginator.Run(98, 1, 100, 10)
	fmt.Printf("CurrentPage: %v\n", pager.CurrentPage)
	fmt.Printf("FirstPage: %v\n", pager.FirstPage)
	fmt.Printf("PrePage: %v\n", pager.PrePage)
	fmt.Printf("Pages: %v\n", pager.Pages)
	fmt.Printf("NextPage: %v\n", pager.NextPage)
	fmt.Printf("LastPage: %v\n", pager.LastPage)
	fmt.Printf("TotalPages: %v\n", pager.TotalPages)
}

```
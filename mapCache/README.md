#### mapCache 缓存工具概述
gotool 提供了基于 sync.map 的缓存操作工具，可操高效地实现数据缓存；

#### 缓存配置
请打开(不存在则创建 ./config.ini)，添加或修改配置
```ini
[MapCache]
# 缓存垃圾回收间隔时间，单位 秒
# 如果缓存数据量教大 [ > 10W ] 应该延长此时间，如 6 小时 6 * 3600
GCIntervalTime=900
```
#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/mapCache"
)
```

#### 使用示例
```go
package main

import (
	"fmt"
	"time"

	"github.com/cnlesscode/gotool/mapCache"
)

type Person struct {
	Name string
	Age  int
}

// 模拟一个动态获取数据函数
func getData(args ...any) (any, error) {
	fmt.Printf("%v\n", "动态函数执行了")
	fmt.Printf("参数传递 : %v\n", args)
	return []Person{{"张三", 18}, {"lisi", 18}}, nil
}

func main() {
	// 第1次获取
	persons := mapCache.Cache("Person", -1, getData, 22)
	fmt.Printf("persons: %T %v\n", persons, persons)

	// 第2次获取
	// 使用断言转换数据类型示例
	time.Sleep(time.Second * 3)
	persons2, ok := mapCache.Cache("Person", -1, getData, 22).([]Person)
	if ok {
		fmt.Printf("persons2: %T %v\n", persons2, persons2[0].Name)
	}

	// 设置缓存
	mapCache.Set("a", 100, "a...")
	mapCache.Set("b", 100, "b...")

	// 删除缓存
	mapCache.Remove("a")

	// 清空缓存
	// mapCache.Clear()

	// 观察缓存 map
	mapCache.MapCacher.Range(func(key, value any) bool {
		fmt.Printf("key: %v\n", key)
		return true
	})
}
```
### 随机数据工具包

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/random"
)
```

#### RangeIntRand(最小值, 最大值 int64) int64
函数功能 : 生成一个指定范围的随机整数
返回格式 : int64
```go
fmt.Printf("%v\n", random.RangeIntRand(1111, 999999))
```

#### RangeFloat(整数范围[]int{}, 小数范围[]int{}) string
函数功能 : 生成一个指定范围的随机小数
返回格式 : float64
```go
floatResult := random.RangeFloat([]int{100, 200}, []int{100, 200})
fmt.Printf("floatResult: %v\n", floatResult)
// 155.198
```

#### RandomString(长度 int, 数字数量 int) string
函数功能 : 生成一个指定长度的随机字符串
返回格式 : string
```go
fmt.Printf("random.RandomString(6, 3): %v\n", random.RandomCharacters(6, 3))
// 5k9nm7
```

#### UUID() string
函数功能 : 生成一个UUID
返回格式 : string
基础工具 : "github.com/google/uuid"
```go
fmt.Printf("%v\n", random.UUID())
// 71042359-59b0-46b1-9a83-3ebfad75114a
```

#### Md5UUID() string
函数功能 : 生成一个基于 UUID 的 Md5 随机字符串
返回格式 : string
```go
fmt.Printf("%v\n", random.Md5UUID())
// df19971fa366e8461223c1dd1d037de0
```

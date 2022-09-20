### 日期时间工具包

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/datetime"
)
```

#### TimeStampToDatatime(时间戳 int64) string 
函数功能 : 时间戳转日期时间
返回格式 : 字符串形式的日期时间
```go
fmt.Printf("%v", datetime.TimeStampToDatatime(1658997290))
// 2022-07-28 16:34:50
```

#### TimeStampToDatatimeSlice(时间戳 int64) []string
函数功能 : 时间戳转日期时间切片形式
返回格式 : 字符串切片
```go
fmt.Printf("%v", datetime.TimeStampToDatatimeSlice(1658997290))
// [2022 07 28 16 34 50]
```

#### DateTimeToTimeStamp(日期时间 string) int64
函数功能 : 日期时间转时间戳
返回格式 : 时间戳 int64
```go
fmt.Printf("%v\n", datetime.DateTimeToTimeStamp("2022-07-28 16:34:50"))
// 1658997290
```

#### FormatPastTime(时间戳 int64) (int, string, string)
函数功能 : 获取过去时间并格式化
返回格式 : 时间差, 过去时间 [ 英文 ], 过去时间 [ 中文 ]
```go
func main() {
	timeDifference, strEn, strZh := datetime.FormatPastTime(1660707752 - 3600)
	fmt.Printf("timeDifference: %v\n", timeDifference)
	fmt.Printf("strEn: %v\n", strEn)
	fmt.Printf("strZh: %v\n", strZh)
}
// 1小时前 || ** 天前 || *** 分钟前 ...
```

#### CountDaysOfAMonth(某年某月 string) (int, error)
函数功能 : 计算某年某月天数
返回格式 : 天数, 错误
```go
func main() {
	days, err := datetime.CountDaysOfAMonth("200402")
	if err == nil {
		fmt.Printf("days: %v\n", days)
	} else {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	// 29
}
```

#### IsLeapYear(年份 int) bool
函数功能 : 判断某年是否为闰年
返回格式 : boolean
```go
fmt.Printf("datetime.IsLeapYear(2000): %v\n", datetime.IsLeapYear(2000)) // true
fmt.Printf("datetime.IsLeapYear(2001): %v\n", datetime.IsLeapYear(2001)) // false
```
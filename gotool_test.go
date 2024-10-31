package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/datetime"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	dt := datetime.New()
	// 初始化当前时间
	dt.Now()
	Loger.Printf("当前时间 : %s", dt.Result)
	// 下一月
	nextMonth := dt.Swicth("month", 1)
	Loger.Printf("一月后时间 : %s", nextMonth.Result)
	nextDay := dt.Swicth("day", 1)
	fmt.Printf("nextDay: %v\n", nextDay)
}

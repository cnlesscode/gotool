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
	fmt.Printf("dt.Result: %v\n", dt.Result)
	// 下一年
	nextYear := dt.SwicthYear(1)
	fmt.Printf("nextYear.Result: %v\n", nextYear.Result)
	// 上一年
	prevYear := dt.SwicthYear(-1)
	fmt.Printf("prevYear.Result: %v\n", prevYear.Result)
	// 未来3个月
	nextMonth := dt.SwicthMonth(3)
	fmt.Printf("nextMonth.Result: %v\n", nextMonth.Result)
	// 上一月
	prevMoth := dt.SwicthMonth(-1)
	fmt.Printf("prevYear.Result: %v\n", prevMoth.Result)
}

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
	for i := 0; i < 19; i++ {
		dtn := dt.Swicth("month", -1*i)
		fmt.Printf("dtn.Result: %v\n", dtn.Result)
	}
}

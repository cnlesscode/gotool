package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/maths"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	var num32 float32 = 12.3
	num32, _ = maths.DecimalPlaces32(num32, 1)
	fmt.Printf("num: %v\n", num32)

	var num64 float64 = 12000008989.31878887
	num64, _ = maths.DecimalPlaces64(num64, 2)
	fmt.Printf("num: %v\n", num64)

	fmt.Printf("num: %v\n", maths.FloatToString(num64, 2))

	fmt.Printf("maths.Floor(float64(num32)): %v\n", maths.Round(float64(num32)))
}

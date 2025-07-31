package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/gmd5"
)

// 单元测试

// 测试命令 : go test -v -run=TestT
func TestT(t *testing.T) {
	res := gmd5.Md5("123456")
	fmt.Printf("res: %v\n", res)
}

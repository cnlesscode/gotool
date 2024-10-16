package gotool

import (
	"fmt"
	"testing"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	fmt.Printf("Root: %v\n", Root)
	fmt.Printf("OS: %v\n", OS)
}

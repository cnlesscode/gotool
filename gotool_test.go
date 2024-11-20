package gotool

import (
	"fmt"
	"testing"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	ip := GetNetworkIP()
	fmt.Printf("ip: %v\n", ip)
}

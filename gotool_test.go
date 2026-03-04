package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/ip"
)

// 单元测试

// 测试命令 : go test -v -run=TestT1
func TestT1(t *testing.T) {
	ip := ip.GetLocalIP()
	fmt.Printf("ip: %v\n", ip)
}

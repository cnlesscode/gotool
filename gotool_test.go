package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/gintool"
)

// 单元测试

// 测试命令 : go test -v -run=TestT
func TestT(t *testing.T) {
	AllowTypes := "image/*,image/png,image/gif"
	fmt.Println(gintool.MatchMimeType(AllowTypes, "image/jpeg"))
}

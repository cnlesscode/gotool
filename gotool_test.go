package gotool

import (
	"testing"
)

// 单元测试

// 测试命令 : go test -v -run=TestT
func TestT(t *testing.T) {
	SetLogLevel(3)
	LogDebug("debug")
	LogError("error")
	LogOk("ok")
	LogFatal("fatal")
}

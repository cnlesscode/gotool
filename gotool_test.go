package gotool

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	// your code
	Loger.Info("hello world")

	for i := 0; i < 10; i++ {
		go func(step int) {
			FileLoger.WithFields(logrus.Fields{"data": step}).Info("hello world")
		}(i)
	}

	time.Sleep(time.Second * 3)
}

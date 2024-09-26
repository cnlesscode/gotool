package gotool

import (
	"os"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var Loger *logrus.Logger
var FileLoger *logrus.Logger

func init() {

	// 命令行直接输出
	Loger = logrus.New()
	Loger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	Loger.Out = colorable.NewColorableStdout()

	// 记录到文件
	FileLoger = logrus.New()
	FileLoger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	file, err := os.OpenFile("./run_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		FileLoger.Out = file
	}
}

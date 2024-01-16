package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cnlesscode/gotool/gfs"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	result := gfs.ScanDirStruct{
		Name:  "根目录",
		Path:  "./a",
		IsDir: true,
	}
	err := gfs.ScanDir(true, &result)
	if err == nil {
		ShowDir(&result, 0)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func ShowDir(root *gfs.ScanDirStruct, step int) {
	fmt.Printf(strings.Repeat("  ", step)+"|_ %s [目录]\n", root.Name)
	step++
	for _, v := range root.Sons {
		if v.IsDir {
			ShowDir(v, step)
		} else {
			fmt.Printf(strings.Repeat("  ", step)+"|_ %s [文件]\n", v.Name)
		}
	}
}

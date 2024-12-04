package gotool

import (
	"fmt"
	"strconv"
	"testing"
)

// 单元测试
// 测试命令 : go test -v -run=TestWrite
func TestWrite(t *testing.T) {
	bf := NewBinaryFile("./data", 100)
	defer bf.Close()

	for i := 1; i <= 300; i++ {
		bf.Write([]byte(strconv.Itoa(i) + "中文:Hello World!"))
	}

}

// 测试命令 : go test -v -run=TestRead
func TestRead(t *testing.T) {
	bf := NewBinaryFile("./data", 100)
	defer bf.Close()

	res, err := bf.Read(299, 8)
	fmt.Printf("res: %s\n", res)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

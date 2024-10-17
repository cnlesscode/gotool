package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/nlp"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	keyWords, keyWordsString := nlp.Cut("我爱我的祖国！")
	fmt.Printf("keyWords: %v\n", keyWords)
	fmt.Printf("keyWordsString: %v\n", keyWordsString)
}

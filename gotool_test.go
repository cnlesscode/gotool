package main

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/gDom"
	"golang.org/x/net/html"
)

// 单元测试
// 测试命令 : go test -v -run=TestMain
func TestMain(t *testing.T) {
	// 获取网页源码，创建 html node
	htmlNode, err := gDom.InitByFile("./demoData/html.html")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 寻找 body 节点
	body, err := gDom.FindNode(htmlNode, html.ElementNode, "body")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 寻找 p 标签 [ 多个 ]
	pNodes := make([]*html.Node, 0)
	gDom.FindNodes(body, html.ElementNode, "p", &pNodes)

	// 将 html 转换为项目
	items := make([]gDom.ItemForContent, 0)
	gDom.TransformHTMLToItems(body, &items)
	fmt.Printf("items: %v\n", items)
}

package gDom

import (
	"errors"
	"regexp"
	"strings"

	"github.com/cnlesscode/gotool/gfs"
	"github.com/cnlesscode/gotool/request"
	"golang.org/x/net/html"
)

type ContentItem struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Html struct {
	HTMLNode *html.Node
	Body     *html.Node
	Items    []ContentItem
}

// 从一个文件初始化 HTML 对象
func InitByFile(filePath string) (*html.Node, error) {
	htmlNode := &html.Node{}
	htmlContent, err := gfs.ReadFile(filePath)
	if err != nil {
		return htmlNode, err
	}
	return InitByContent(htmlContent)
}

// 从内容初始化 HTML 对象
func InitByContent(htmlContent string) (*html.Node, error) {
	htmlNode := &html.Node{}
	reg := regexp.MustCompile("[\n\t\r]")
	htmlContent = reg.ReplaceAllString(htmlContent, "")
	reg = regexp.MustCompile("(<br />)|(<br>)")
	htmlContent = reg.ReplaceAllString(htmlContent, "\n")
	htmlContent = strings.Replace(htmlContent, "  ", "", -1)
	superNodes, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return htmlNode, errors.New("无法解析HTML")
	}
	for htmlNode = superNodes.FirstChild; htmlNode != nil; htmlNode = htmlNode.NextSibling {
		if htmlNode.Type == 3 && htmlNode.Data == "html" {
			return htmlNode, nil
		}
	}
	return htmlNode, errors.New("无法解析HTML")
}

// 从 url 初始化 HTML 对象
func InitByUrl(url string) (*html.Node, error) {
	htmlNode := &html.Node{}
	var err error
	res, err := request.GET(url, nil, nil)
	if err != nil {
		return htmlNode, err
	}
	return InitByContent(res)
}

/*
NodeType :
0 ErrorNode NodeType
1 TextNode
2 DocumentNode
3 ElementNode
4 CommentNode
5 DoctypeNode
*/

// 查找单一节点
func FindNode(node *html.Node, nodeType html.NodeType, tagName string) (*html.Node, error) {
	for nodeIn := node.FirstChild; nodeIn != nil; nodeIn = nodeIn.NextSibling {
		if nodeIn.Type == nodeType && nodeIn.Data == tagName {
			return nodeIn, nil
		} else {
			FindNode(nodeIn, nodeType, tagName)
		}
	}
	return node, errors.New("not found")
}

// 查找对应标签的全部节点
func FindNodes(node *html.Node, nodeType html.NodeType, tagName string, nodes *[]*html.Node) {
	for nodeIn := node.FirstChild; nodeIn != nil; nodeIn = nodeIn.NextSibling {
		if nodeIn.Type == nodeType && nodeIn.Data == tagName {
			*nodes = append(*nodes, nodeIn)
		} else {
			FindNodes(nodeIn, nodeType, tagName, nodes)
		}
	}
}

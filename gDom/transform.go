package gDom

import (
	"strings"

	"golang.org/x/net/html"
)

type ItemForContent struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// 将项目转化为html形式
func TransformItemsToHtml(items []ItemForContent) string {
	var htmlContent string = ""
	for _, item := range items {
		if item.Type == "a" {
			var linkArray = strings.Split(item.Content, "](")
			htmlContent += `<p><a href="` + linkArray[1][0:len(linkArray[1])-1] + `" target="_blank">` + linkArray[0][1:] + `</a><p>`
		} else if item.Type == "hr" {
			htmlContent += "<p><hr /></p>"
		} else if item.Type == "img" || item.Type == "image" {
			htmlContent += `<img src="` + item.Content + `" />`
		} else if item.Type == "video" || item.Type == "audio" {
			htmlContent += `<${item.Type} src="` + item.Content + `"></${item.Type}>`
		} else if item.Type == "p" {
			htmlContent += "<p>" + item.Content + "</p>"
		} else {
			htmlContent += "<p><" + item.Type + ">" + item.Content + "<" + item.Type + "/></p>"
		}
	}
	return htmlContent
}

// 将 html 源码转换为项目数组
func TransformHTMLToItems(parentNode *html.Node, items *[]ItemForContent) {
	for node := parentNode.FirstChild; node != nil; node = node.NextSibling {
		if node.Type > 1 {
			if node.FirstChild == nil {
				node.AppendChild(&html.Node{
					Type: html.TextNode,
					Data: "",
				})
			}
			TransformHTMLToItems(node, items)
		} else {
			AnalysisTag(node, items)
		}
	}
}

// 解析标签
func AnalysisTag(node *html.Node, items *[]ItemForContent) {
	tagData := node.Data
	parentNode := node.Parent
	tagType := parentNode.Data
	// 文本
	if tagType == "hr" {
		*items = append(*items, ItemForContent{Type: "hr", Content: "..."})
	} else if tagType == "a" {
		href := Href(parentNode)
		*items = append(
			*items,
			ItemForContent{
				Type:    "a",
				Content: "[" + tagData + "](" + href + ")",
			},
		)
	} else if tagType == "img" || tagType == "video" || tagType == "audio" {
		src := Src(parentNode)
		*items = append(
			*items,
			ItemForContent{
				Type:    tagType,
				Content: src,
			},
		)
	} else {
		if tagData != "" {
			*items = append(*items, ItemForContent{Type: tagType, Content: tagData})
		}
	}
}

// 获取 href 属性
func Href(node *html.Node) string {
	href := ""
	for _, item := range node.Attr {
		key := strings.ToLower(item.Key)
		if key == "href" {
			href = item.Val
			break
		}
	}
	return href
}

// 获取 src 属性
func Src(node *html.Node) string {
	src := ""
	for _, item := range node.Attr {
		key := strings.ToLower(item.Key)
		if key == "src" {
			src = item.Val
			break
		}
	}
	return src
}

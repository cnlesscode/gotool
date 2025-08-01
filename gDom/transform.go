package gDom

import (
	"strings"

	"golang.org/x/net/html"
)

// 将项目转化为html形式
func TransformItemsToHtml(items []ContentItem) string {
	var htmlContent string = ""
	for _, item := range items {
		switch item.Type {
		case "a":
			var linkArray = strings.Split(item.Content, "](")
			htmlContent += `<p><a href="` + linkArray[1][0:len(linkArray[1])-1] + `" target="_blank">` + linkArray[0][1:] + `</a><p>`
		case "hr":
			htmlContent += "<p><hr /></p>"
		case "img", "image":
			htmlContent += `<img src="` + item.Content + `" />`
		case "video", "audio":
			htmlContent += `<${item.Type} src="` + item.Content + `"></${item.Type}>`
		case "p":
			htmlContent += "<p>" + item.Content + "</p>"
		default:
			htmlContent += "<p><" + item.Type + ">" + item.Content + "<" + item.Type + "/></p>"
		}
	}
	return htmlContent
}

// 将 html 源码转换为项目数组
func TransformHTMLToItems(parentNode *html.Node, items *[]ContentItem) {
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
func AnalysisTag(node *html.Node, items *[]ContentItem) {
	tagData := node.Data
	parentNode := node.Parent
	tagType := parentNode.Data
	// 文本
	switch tagType {
	case "hr":
		*items = append(*items, ContentItem{Type: "hr", Content: "..."})
	case "a":
		href := Href(parentNode)
		*items = append(
			*items,
			ContentItem{
				Type:    "a",
				Content: "[" + tagData + "](" + href + ")",
			},
		)
	case "img", "video", "audio":
		src := Src(parentNode)
		*items = append(
			*items,
			ContentItem{
				Type:    tagType,
				Content: src,
			},
		)
	default:
		if tagData != "" {
			*items = append(*items, ContentItem{Type: tagType, Content: tagData})
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

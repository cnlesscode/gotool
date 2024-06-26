package gintool

import (
	"bytes"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

// POST 过滤
func SafePOST(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		return
	}
	var specialChars = []string{"<", ">", "%3C", "%3E"}
	var specialCharsTo = []string{"_", "_", "_", "_"}
	for idx, char := range specialChars {
		body = bytes.ReplaceAll(body, []byte(char), []byte(specialCharsTo[idx]))
	}
	// 将过滤后的请求体设置回请求中
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	ctx.Request.ContentLength = int64(len(body))
}

// GET 过滤
func SafeQuery(ctx *gin.Context, key string) string {
	data := ctx.Query(key)
	if data == "" {
		return data
	}
	data = strings.ReplaceAll(data, "<", "&lt;")
	data = strings.ReplaceAll(data, ">", "&gt;")
	return data
}

// 字符串 过滤
func SafeData(data string) string {
	data = strings.ReplaceAll(data, "<", "&lt;")
	data = strings.ReplaceAll(data, ">", "&gt;")
	return data
}

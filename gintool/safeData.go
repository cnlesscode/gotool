package gintool

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// POST 过滤
func SafePOST(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	bodyString := string(body)
	var specialChars = []string{"<", "%3C", "%3c", ">", "%3E", "%3e"}
	var specialCharsTo = []string{"‹", "‹", "‹", "›", "›", "›"}
	for idx, char := range specialChars {
		bodyString = strings.ReplaceAll(bodyString, char, specialCharsTo[idx])
	}
	body = []byte(bodyString)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	ctx.Request.ContentLength = int64(len(body))
}

// GET 过滤
func SafeQuery(ctx *gin.Context, key string) string {
	data := ctx.Query(key)
	if data == "" {
		return data
	}
	data = strings.ReplaceAll(data, "<", "‹")
	data = strings.ReplaceAll(data, ">", "›")
	return data
}

// 字符串 过滤
func SafeData(data string) string {
	data = strings.ReplaceAll(data, "<", "‹")
	data = strings.ReplaceAll(data, ">", "›")
	return data
}

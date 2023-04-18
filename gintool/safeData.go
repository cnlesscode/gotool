package gintool

import (
	"bytes"
	"io"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func SafePOST(ctx *gin.Context, newString string, needReplaceString ...string) {
	if ctx.Request.Method == "POST" {
		// 读取请求体中的POST数据
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Abort()
			return
		}
		bodyString, err := url.QueryUnescape(string(bodyBytes))
		if err != nil {
			ctx.Abort()
			return
		}
		for _, v := range needReplaceString {
			bodyString = strings.ReplaceAll(bodyString, v, newString)
			bodyString = strings.ReplaceAll(bodyString, v, newString)
		}
		ctx.Request.Body = io.NopCloser(bytes.NewReader([]byte(bodyString)))
	}
}

func SafeQuery(ctx *gin.Context, newString string, needReplaceString ...string) {
	if ctx.Request.URL.RawQuery == "" {
		return
	}
	rawQuery, err := url.QueryUnescape(string(ctx.Request.URL.RawQuery))
	if err != nil {
		ctx.Abort()
		return
	}

	for _, v := range needReplaceString {
		rawQuery = strings.ReplaceAll(rawQuery, v, newString)
	}
	ctx.Request.URL.RawQuery = rawQuery
}

func SafeData(data string, newString string, needReplaceString ...string) string {
	for _, v := range needReplaceString {
		data = strings.ReplaceAll(data, v, newString)
	}
	return data
}

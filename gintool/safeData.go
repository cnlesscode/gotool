package gintool

import (
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// 请求过滤
func SafeRequest(ctx *gin.Context) {
	// POST
	if ctx.Request.Method == "POST" {
		SafePOST(ctx)
	}
	// GET
	SafeQuery(ctx)
}

// POST 过滤
func SafePOST(ctx *gin.Context) {
	ctx.Request.ParseForm()
	formMap := ctx.Request.PostForm
	for k, item := range formMap {
		safeString := strings.ReplaceAll(item[0], "<", "&lt;")
		safeString = strings.ReplaceAll(safeString, ">", "&gt;")
		ctx.Request.Form.Set(k, safeString)
	}
}

// GET 过滤
func SafeQuery(ctx *gin.Context) {
	if ctx.Request.URL.RawQuery == "" {
		return
	}
	rawQuery, err := url.QueryUnescape(string(ctx.Request.URL.RawQuery))
	if err != nil {
		return
	}
	rawQuery = strings.ReplaceAll(rawQuery, "<", "&lt;")
	rawQuery = strings.ReplaceAll(rawQuery, ">", "&gt;")
	ctx.Request.URL.RawQuery = rawQuery
}

// 字符串 过滤
func SafeData(data string) string {
	data = strings.ReplaceAll(data, "<", "&lt;")
	data = strings.ReplaceAll(data, ">", "&gt;")
	return data
}

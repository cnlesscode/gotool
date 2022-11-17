package gintool

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// 小数型 float64 url 参数获取
// ?id=123 返回 123, true
// ?id=a 返回 0, false
func QueryFloat64(ctx *gin.Context, key string) (float64, bool) {
	val := ctx.Query(key)
	if val == "" {
		return 0, false
	}
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, false
	}
	return valFloat, true
}

// 整数型 int url 参数获取
// ?id=123 返回 123, true
// ?id=a 返回 0, false
func QueryInt(ctx *gin.Context, key string) (int, bool) {
	val := ctx.Query(key)
	if val == "" {
		return 0, false
	}
	varInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return varInt, true
}

// 整数型 int64 url 参数获取
// ?id=123 返回 123, true
// ?id=a 返回 0, false
func QueryInt64(ctx *gin.Context, key string) (int64, bool) {
	val, ok := QueryInt(ctx, key)
	if ok {
		return int64(val), true
	} else {
		return 0, false
	}
}

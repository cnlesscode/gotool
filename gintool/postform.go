package gintool

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// 整数型 int POST 数据获取
func PostFormInt(ctx *gin.Context, key string) (int, bool) {
	val := ctx.PostForm(key)
	if val == "" {
		return 0, false
	}
	varInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return varInt, true
}

// 整数型 int64 POST 数据获取
func PostFormInt64(ctx *gin.Context, key string) (int64, bool) {
	val, ok := PostFormInt(ctx, key)
	if ok {
		return int64(val), true
	} else {
		return 0, false
	}
}

// float64 浮点型 POST 数据获取
func PostFormFloat64(ctx *gin.Context, key string) (float64, bool) {
	val := ctx.PostForm(key)
	if val == "" {
		return 0, false
	}
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, false
	}
	return valFloat, true
}

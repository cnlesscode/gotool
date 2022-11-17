package gintool

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// 整数型 int Param 参数获取
// 路由定义 /:id 访问 /123 返回 123, true
// 路由定义 /:id 访问 /abc 返回 0, false
func ParamInt(ctx *gin.Context, key string) (int, bool) {
	val := ctx.Param(key)
	if val == "" {
		return 0, false
	}
	varInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return varInt, true
}

// 整数型 int64 Param 参数获取
// 路由定义 /:id 访问 /123 返回 123, true
// 路由定义 /:id 访问 /abc 返回 0, false
func ParamInt64(ctx *gin.Context, key string) (int64, bool) {
	val, ok := ParamInt(ctx, key)
	if ok {
		return int64(val), true
	} else {
		return 0, false
	}
}

// 小数型 float64 路由参数获取
// ?id=123 返回 123, true
// ?id=a 返回 0, false
func ParamFloat64(ctx *gin.Context, key string) (float64, bool) {
	val := ctx.Param(key)
	if val == "" {
		return 0, false
	}
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, false
	}
	return valFloat, true
}

package main

import (
	"github.com/cnlesscode/gotool/gintool"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 测试路由
	r.GET("/:name/:age", func(ctx *gin.Context) {
		gintool.Download(ctx, "README.md", "newname", false)
	})

	// 监听指定端口
	r.Run(":80")

}

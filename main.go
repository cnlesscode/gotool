package main

import (
	"fmt"
	"html/template"

	"github.com/cnlesscode/gotool/gintool"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 处理上传
	r.POST("/upload", func(ctx *gin.Context) {
		uploader := &gintool.Upload{}
		// 文件域名称
		uploader.FileName = "file"
		// 文件大小限制, 单位 M
		uploader.MaxSize = 10
		// 允许的类型
		uploader.AllowTypes = "image/jpg,image/jpeg,image/png,image/gif"
		// 允许的扩展名
		uploader.AllowExeNames = "jpg,gif,png"
		// 目标文件夹
		uploader.TargetDir = "./staticg/images/"
		// 文件夹命名规则
		uploader.DirNamingRule = "month"
		// 文件命名规则
		uploader.FileNamingRule = "random"

		err, filePath := uploader.Run(ctx)
		if err == nil {
			ctx.JSON(200, gin.H{"filePath": filePath})
		} else {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(200, gin.H{"error": err.Error()})
		}
	})

	// 展示上传视图
	r.GET("/", func(ctx *gin.Context) {
		t, _ := template.ParseFiles("./templates/upload.html")
		t.Execute(ctx.Writer, gin.H{})
	})

	// 监听指定端口
	r.Run(":80")

}

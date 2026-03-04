package gotool

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/cnlesscode/gotool/cloud/static"
	"github.com/gin-gonic/gin"
)

// 单元测试

// 测试命令 : go test -v -run=TestRunAHttpServer
func TestRunAHttpServer(t *testing.T) {
	// 模拟一个自定义云服务器
	r := gin.Default()
	r.POST("/StaticFileAPI", func(c *gin.Context) {

		// 验证密钥
		secret := c.PostForm("Secret")
		if secret != "123456" {
			c.JSON(0, gin.H{
				"ErrorCode": 400501,
				"Data":      "Secret is error"})
			return
		}

		filePath := c.PostForm("FilePath")
		if filePath == "" {
			c.JSON(0, gin.H{
				"ErrorCode": "400401",
				"Data":      "FilePath is empty"})
			return
		}

		action := c.PostForm("Action")
		switch action {

		// 删除文件
		case "Remove":
			if filePath == "" {
				c.JSON(0, gin.H{
					"ErrorCode": "400401",
					"Data":      "FilePath is empty"})
				return
			}
			os.Remove(filePath)
			c.JSON(0,
				gin.H{
					"ErrorCode": "0",
					"Data":      "ok"})
			return

		// 上传文件
		case "UploadFile":
			sourceFile, err := c.FormFile("FileData")
			if err != nil {
				c.JSON(0, gin.H{
					"ErrorCode": "400501",
					"Data":      "上传文件数据错误"})
				return
			}

			// 确保目标目录存在
			dir := filepath.Dir(filePath)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, os.ModePerm); err != nil {
					c.JSON(0, gin.H{
						"ErrorCode": "400502",
						"Data":      fmt.Sprintf("create dir failed: %v", err)})
					return
				}
			}

			// 保存文件
			err = c.SaveUploadedFile(sourceFile, filePath)
			if err != nil {
				c.JSON(0,
					gin.H{
						"ErrorCode": "40050",
						"Data":      err.Error()})
				return
			}

			c.JSON(0,
				gin.H{
					"ErrorCode": "0",
					"Data":      "ok"})
			return

		// 移动文件
		case "MoveFile":
			targetFile := c.PostForm("TargetFileUrl")
			fmt.Printf("targetFile: %v\n", targetFile)
			fmt.Printf("filePath: %v\n", filePath)
			if targetFile == "" {
				c.JSON(0, gin.H{
					"ErrorCode": "400401",
					"Data":      "TargetFileUrl is empty"})
				return
			}
			dir := filepath.Dir(targetFile)
			fmt.Printf("dir: %v\n", dir)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, os.ModePerm); err != nil {
					c.JSON(0, gin.H{
						"ErrorCode": "400502",
						"Data":      fmt.Sprintf("create dir failed: %v", err)})
					return
				}
			}
			err := os.Rename(filePath, targetFile)
			if err != nil {
				c.JSON(0,
					gin.H{
						"ErrorCode": "40050",
						"Data":      err.Error()})
				return
			}
			c.JSON(0,
				gin.H{
					"ErrorCode": "0",
					"Data":      "ok"})
			return
		// default
		default:
			return
		}

	})
	r.Run(":8080")

}

// 测试命令 : go test -v -run=TestT
func TestT(t *testing.T) {
	staticCloudConfig := static.Config{
		Type:     "Customize",
		Endpoint: "http://192.168.31.189:8080/StaticFileAPI",
		KeyId:    "******",
		Secret:   "123456",
		BaseUrl:  "http://192.168.31.189:8080",
	}
	staticTool := static.New(staticCloudConfig)

	// upload
	err := staticTool.UploadFile("./static/2.txt")
	fmt.Printf("UploadFile err: %v\n", err)

	// Move File
	fileUrl, err := staticTool.MoveFile("./static/1.txt", "./static/2.txt")
	fmt.Printf("fileUrl: %v\n", fileUrl)
	fmt.Printf("MoveFile err: %v\n", err)
}

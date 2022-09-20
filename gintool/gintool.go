package gintool

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/cnlesscode/gotool/gfs"
	"github.com/gin-gonic/gin"
)

// Make A Downloader
func Download(ctx *gin.Context, fileUrl string, remove bool) {
	filename := gfs.GetFileName(fileUrl, "/")
	rw := ctx.Writer
	header := rw.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+filename)
	//使用ioutil包读取文件
	b, err := ioutil.ReadFile(fileUrl)
	if err == nil {
		//写入到响应流中
		rw.Write(b)
	}
	if remove {
		os.Remove(fileUrl)
	}
}

/*
初始化页码及每页展示数量
*/
func QueryPagerInit(ctx *gin.Context) (int, int, int) {
	// 页码
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	// 每页展示数量
	pq, err := strconv.Atoi(ctx.Query("pq"))
	if err != nil {
		pq = 10
	}
	if pq <= 0 {
		pq = 10
	}
	// 分页开始及结束
	limitStart := (page - 1) * pq
	return page, pq, limitStart
}

// 初始化指定的 url 数据为整数
func QueryInt(ctx *gin.Context, name string) int {
	val, err := strconv.Atoi(ctx.Query(name))
	if err != nil {
		return 0
	}
	return val
}

// 初始化指定的 ParamInt 数据为整数
func ParamInt(ctx *gin.Context, name string, defaultVal int) int {
	val, err := strconv.Atoi(ctx.Param(name))
	if err != nil {
		return defaultVal
	}
	return val
}

// 初始化 Actions
func Actions(ctx *gin.Context) []string {
	action := ctx.Param("page")
	if action == "" {
		return nil
	}
	// 截取 /
	action = strings.Trim(action, "/")
	// 分割
	return strings.Split(action, "/")
}

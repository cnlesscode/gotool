package gimage

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/cnlesscode/gotool/gfs"
)

// 将 Base64 转换为图片
func Base64ToImage(data string, targetDir string, imageName string) (string, error) {
	reg := regexp.MustCompile(`(?U)data:image\/(.*);base64,`)
	res := reg.FindStringSubmatch(data)
	if len(res) < 2 {
		return "", errors.New("base64 data error")
	}
	if !gfs.PathExists(targetDir) {
		os.MkdirAll(targetDir, 0777)
	}
	imageName = targetDir + imageName
	imageName = imageName + "." + res[1]
	data = strings.Replace(data, res[0], "", -1)
	bt, _ := base64.StdEncoding.DecodeString(data)
	err := os.WriteFile(imageName, bt, 0777)
	if err != nil {
		return "", err
	} else {
		return imageName, nil
	}
}

// 图片转Base64
func ImageToBase64(imageUri string, removeImage bool) (string, error) {
	file, err := os.Open(imageUri)
	if err != nil {
		return "", err
	}
	imgByte, err := io.ReadAll(file)
	file.Close()
	if err != nil {
		return "", err
	}
	mimeType := http.DetectContentType(imgByte)
	baseImg := "data:" + mimeType + ";base64," + base64.StdEncoding.EncodeToString(imgByte)
	if removeImage {
		os.Remove(imageUri)
	}
	return baseImg, nil
}

package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var TimeoutSecond int = 15

// 基础
func Base(method string, uri string, data map[string]string, headers map[string]string) (string, error) {
	urlData := url.Values{}
	for k, v := range data {
		urlData.Set(k, v)
	}
	if method == "GET" || method == "DELETE" {
		uri += "?" + urlData.Encode()
	}
	req, err := http.NewRequest(method, uri, strings.NewReader(urlData.Encode()))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if method == "POST" || method == "PUT" {
		ContentType, ok := headers["Content-Type"]
		if !ok || ContentType == "" {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	if err != nil {
		return "", err
	}
	client := &http.Client{Timeout: time.Second * time.Duration(TimeoutSecond)}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respdata), nil
}

// GET 请求
func GET(uri string, data map[string]string, headers map[string]string) (string, error) {
	return Base("GET", uri, data, headers)
}

// POST 请求
func POST(uri string, data map[string]string, headers map[string]string) (string, error) {
	return Base("POST", uri, data, headers)
}

// PUT 请求
func PUT(uri string, data map[string]string, headers map[string]string) (string, error) {
	return Base("PUT", uri, data, headers)
}

// PUT 请求
func DELETE(uri string, data map[string]string, headers map[string]string) (string, error) {
	return Base("DELETE", uri, data, headers)
}

// 上传文件
func UploadFile(uri string, filePath string, postData map[string]string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file %s: %w", filePath, err)
	}
	defer file.Close()

	// 获取文件名
	filename := filepath.Base(filePath)

	// 创建 multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 写入普通字段
	for k, v := range postData {
		_ = writer.WriteField(k, v)
	}
	_ = writer.WriteField("FilePath", filePath)

	// 写入文件字段（字段名假设为 "FileData"，需与服务端约定）
	part, err := writer.CreateFormFile("FileData", filename)
	if err != nil {
		return fmt.Errorf("create form file: %w", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("copy file to form: %w", err)
	}

	// 关闭 writer，必须在发送前调用
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("close multipart writer: %w", err)
	}

	// 发送请求
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var res map[string]string = map[string]string{}
	err = json.Unmarshal(respdata, &res)
	if err != nil {
		return err
	}
	if res["ErrorCode"] == "0" {
		return nil
	}

	return errors.New(res["Data"])
}

package request

import (
	"io"
	"net/http"
	"net/url"
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

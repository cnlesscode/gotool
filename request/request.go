package request

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// 基础
func Base(method string, uri string, data map[string]string, headers map[string]string) (*http.Request, error) {
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
	return req, err
}

// GET 请求
func GET(uri string, data map[string]string, headers map[string]string) (string, error) {
	req, err := Base("GET", uri, data, headers)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
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

// POST 请求
func POST(uri string, data map[string]string, headers map[string]string) (string, error) {
	req, err := Base("POST", uri, data, headers)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
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

// PUT 请求
func PUT(uri string, data map[string]string, headers map[string]string) (string, error) {
	req, err := Base("PUT", uri, data, headers)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
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

// PUT 请求
func DELETE(uri string, data map[string]string, headers map[string]string) (string, error) {
	req, err := Base("DELETE", uri, data, headers)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
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

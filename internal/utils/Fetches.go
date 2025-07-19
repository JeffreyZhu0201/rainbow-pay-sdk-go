/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-05-07 14:58:52
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-05-09 19:35:47
 * @FilePath: \RocketVPN\go-backend\utils\Fetches.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// 传入目标url和参数map，返回post请求的结果
func FetchPost(targetUrl string, params map[string]any) (string, error) {

	//将map转为json格式
	jsonData, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(targetUrl, "application/json", bytes.NewBuffer(jsonData)) //将json转为字符串
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// 传入目标url和参数map，返回post请求的结果(使用form表单)
func FetchPostForm(targetUrl string, params map[string]any) (string, error) {
	formData := url.Values{}
	for key, value := range params {
		formData.Add(key, fmt.Sprintf("%v", value))
	}
	log.Println("")

	resp, err := http.PostForm(targetUrl, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// 传入目标url和参数map，返回post请求的结果(使用query参数)
func FetchPostWithQuery(targetUrl string, params map[string]any) (string, error) {
	// Build query string
	query := "?"
	for key, value := range params {
		query += fmt.Sprintf("%s=%v&", key, value)
	}
	// Remove last &
	query = query[:len(query)-1]

	// Append query to URL
	fullUrl := targetUrl + query

	resp, err := http.Post(fullUrl, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(fullUrl))
	// Parse response as JSON
	return string(body), nil
}

func Post(targetUrl string) (string, error) {
	// Create a new request
	req, err := http.NewRequest("POST", targetUrl, nil)
	if err != nil {
		return "", err
	}

	// Add headers
	req.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Host", "api.xn--6krw8b915a62m.xn--io0a7i")
	req.Header.Set("Connection", "keep-alive")

	// Create HTTP client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("URL:", targetUrl)
	log.Println("Response:", string(body))

	return string(body), nil
}

// 传入目标url和参数map，返回get请求的结果
func FetchGet(targetUrl string, params map[string]any) (string, error) {
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

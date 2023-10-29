package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
	Message      Message `json:"message"`
}

type OpenAIResponse struct {
	Choices []Choice       `json:"choices"`
	Created float64        `json:"created"`
	ID      string         `json:"id"`
	Model   string         `json:"model"`
	Object  string         `json:"object"`
	Usage   map[string]int `json:"usage"`
}

func Openai(apiKey string, q string) (*OpenAIResponse, error) {
	// 设置 API 端点 URL 和你的 API 密钥
	endpointURL := "https://api.openai.com/v1/chat/completions"

	// 构造请求数据
	requestData := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{ // 修改此处的数据类型
			{"role": "user", "content": q},
		},
		"temperature": 0.7,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("无法编码请求数据为 JSON:", err)
		return nil, err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", endpointURL, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		fmt.Println("创建 HTTP 请求时出错:", err)
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("执行 HTTP 请求时出错:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取和解析响应
	var responseData OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("解析响应时出错:", err)
		return nil, err
	}

	// 返回解析后的响应数据
	return &responseData, nil
}

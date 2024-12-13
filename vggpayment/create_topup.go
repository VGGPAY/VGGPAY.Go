// Package vggpayment file: vggpayment/create_topup.go

package vggpayment

import (
	"encoding/json"
	"fmt"
)

// CreateTopUp 创建充值订单并发送请求
func CreateTopUp(config *AuthConfigConfig, data map[string]interface{}) (int, string, error) {
	// 从config中提取项目 ID, SecretKey, SecretIV
	projectId := config.ProjectId
	SecretKey := config.SecretKey
	SecretIV := config.SecretIV

	// 给 data 设置默认值（确保必要字段存在）
	if data["projectid"] == nil {
		data["projectid"] = projectId
	}
	if data["firewall"] == nil {
		data["firewall"] = "2"
	}

	// 将数据转换为JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal data: %v", err)
	}

	// 加密请求数据
	encryptedData := EncryptData(string(jsonData), SecretIV, SecretKey)

	// 构建POST数据
	postData := map[string]interface{}{
		"data":      encryptedData,
		"projectid": projectId,
	}
	postDataBytes, err := json.Marshal(postData)
	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal post data: %v", err)
	}

	// 发送HTTP请求
	url := "https://sapi.vggpay.com/api/v2/createtopup"
	statusCode, responseBody, err := SendRequest(url, postDataBytes)
	if err != nil {
		return statusCode, responseBody, fmt.Errorf("request failed: %v", err)
	}

	// 返回状态码和响应体
	return statusCode, responseBody, nil
}

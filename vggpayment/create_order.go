// Package vggpayment file: vggpayment/create_order.go

package vggpayment

import (
	"encoding/json"
	"fmt"
)

// CreateOrder Create an order and send a request
func CreateOrder(config *AuthConfigConfig, data map[string]interface{}) (int, string, error) {
	// 从config中提取项目 ID, SecretKey, SecretIV
	projectId := config.ProjectId
	SecretKey := config.SecretKey
	SecretIV := config.SecretIV

	// Set default values  (make sure required fields exist)
	if data["projectid"] == nil {
		data["projectid"] = projectId
	}
	if data["firewall"] == nil {
		data["firewall"] = "2"
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal data: %v", err)
	}

	// Encrypting request data
	encryptedData := EncryptData(string(jsonData), SecretIV, SecretKey)

	// Constructing POST data
	postData := map[string]interface{}{
		"data":      encryptedData,
		"projectid": projectId,
	}
	postDataBytes, err := json.Marshal(postData)
	if err != nil {
		return 0, "", fmt.Errorf("failed to marshal post data: %v", err)
	}

	// Sending HTTP Request
	url := "https://sapi.vggpay.com/api/v2/createorder"
	statusCode, responseBody, err := SendRequest(url, postDataBytes)
	if err != nil {
		return statusCode, responseBody, fmt.Errorf("request failed: %v", err)
	}

	// Return status code and response body
	return statusCode, responseBody, nil
}

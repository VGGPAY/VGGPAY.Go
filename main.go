package main

import (
	"VGGPAY.Go/vggpayment"

	"fmt"
	"log"
)

func main() {

	// Set your project ID and key information
	config := vggpayment.AuthConfig(
		"999DEMO",
		"88d4012da55e249ab48cffbe2f19d6326e524680d5dfa8b5990b02fdc9473682",
		"6ad4dabbb9844769fb33e8655a78a7fc")

	fmt.Println("Create top up example:")

	data1 := map[string]interface{}{
		"m_userid":   "userdemo001",
		"firewall":   "2",
		"notify_url": "https://my-notify-api.com",
	}

	statusCode1, responseBody1, err1 := vggpayment.CreateTopUp(config, data1)
	if err1 != nil {
		log.Fatal("Error:", err1)
	} else {
		fmt.Println("Response Status:", statusCode1)
		fmt.Println("Response Body:", responseBody1)
	}

	fmt.Println("Create order example:")

	data2 := map[string]interface{}{
		"m_orderid":    "yourShopOrder12345679",
		"currency":     "EUR",
		"amount":       "815.23",
		"notify_url":   "https://my-notify-api.com",
		"notify_txt":   "{\"Product\":\"iPhone 13\",\"modelColor\":\"red\",\"myStrings\":\"Custom Strings\"}",
		"time_out":     "1200",
		"redirect_url": "",
		"firewall":     "2",
	}

	statusCode2, responseBody2, err2 := vggpayment.CreateOrder(config, data2)
	if err2 != nil {
		log.Fatal("Error:", err2)
	} else {
		fmt.Println("Response Status:", statusCode2)
		fmt.Println("Response Body:", responseBody2)
	}

	fmt.Println("VGGPay sends callback data to your server. Here is an example of how to decrypt the callback data.:")

	encryptedData := "jdtdW1+nP8D3geSHQ0+5h0V5Dpez3Lmon0dpW6Dd4BnOEPDqdWNeuow7MM0XHxshHDJxP1QXslO81Enw+JryoRqEWCQYaS282TjqxXtxXfkL1NeLqwlJsKk/EKCnlDGZy2tP5fgCrGaWxAhGWVUtrTppidgFmJrHGh1c5qKFZe0jsxzIY+YI37KyhsOCJzKYgP4GwkQljh+SNF0AuH6vmnI710cczfIXjT2/GJjnJugnVtYuV/W4UN8qgPj3NAZWuDXM6oe1xTufGb8lNU1HctbBRheqUU2/xyGqJz8AOZnb9Z6//r7U90vfdhBolZ94PojBucyifPyShnTaNS+Uy4ZB6UmACWmFtDZTjmOzLbm/dL0ppFVxqMbxQjpTr7OeZKHEkMZLoxJygjThoTbUNQspM5DeVgwgeXHBUnlGma9MkOuIPppfbWrGbtlpZVj6"

	decryptedData, err3 := vggpayment.DecryptData(config, encryptedData)
	if err3 != nil {
		log.Fatal("Error decrypting data:", err3)
	}

	fmt.Println("Decrypted Data:", decryptedData)

}

// Package vggpayment file:vggpayment/requests.go

package vggpayment

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

// SendRequest Sending HTTP Request
func SendRequest(url string, data []byte) (int, string, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	// Sending a POST request
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return 0, "", fmt.Errorf("request failed: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Failed to close response body:", err)
		}
	}()

	// Determine if the HTTP response status code is 200
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	// Reading the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Return status code and response body
	return resp.StatusCode, string(responseBody), nil
}

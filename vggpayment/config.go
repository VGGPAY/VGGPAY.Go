// Package vggpayment file:vggpayment/config.go
package vggpayment

// AuthConfigConfig Configuration Structure

type AuthConfigConfig struct {
	ProjectId string
	SecretKey string
	SecretIV  string
}

// AuthConfig Create a new configuration object
func AuthConfig(projectId, secretKey, secretIV string) *AuthConfigConfig {
	return &AuthConfigConfig{
		ProjectId: projectId,
		SecretKey: secretKey,
		SecretIV:  secretIV,
	}
}

// Package vggpayment file: vggpayment/decryptdata.go

package vggpayment

func DecryptData(config *AuthConfigConfig, data string) (string, error) {
	SecretKey := config.SecretKey
	SecretIV := config.SecretIV
	decryptedData, err := decryptAES(data, SecretKey, SecretIV)
	if err != nil {
		return "", err
	}
	return decryptedData, nil
}

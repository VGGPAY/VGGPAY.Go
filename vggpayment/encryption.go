// Package vggpayment file:vggpayment/encryption.go

package vggpayment

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"log"
)

// EncryptData AES encryption function
func EncryptData(plaintext string, SecretIV string, SecretKey string) string {
	iv := hex2bin(SecretIV)
	key := hex2bin(SecretKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	paddedData := pad([]byte(plaintext), aes.BlockSize)
	encrypter := cipher.NewCBCEncrypter(block, iv)

	encrypted := make([]byte, len(paddedData))
	encrypter.CryptBlocks(encrypted, paddedData)

	return base64.StdEncoding.EncodeToString(encrypted)
}

// Convert hex string to byte array
func hex2bin(hexStr string) []byte {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// Pad the data to make it a multiple of 16
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 实现 PKCS7 反向填充算法
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func decryptAES(encryptedData string, key string, iv string) (string, error) {
	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	ivBytes, err := hex.DecodeString(iv)
	if err != nil {
		return "", err
	}

	encrypted, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, ivBytes)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	decrypted = pkcs7Unpad(decrypted)

	return string(decrypted), nil
}

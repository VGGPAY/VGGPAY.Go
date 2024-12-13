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

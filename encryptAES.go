package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func encryptAES(input, key string) string {
	// Convert key to []byte
	keyBytes := []byte(key)

	// Create AES block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return ""
	}

	// Add padding to the input text, if necessary
	inputBytes := []byte(input)
	padding := aes.BlockSize - len(inputBytes)%aes.BlockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	inputBytes = append(inputBytes, paddingText...)

	// We create the cipher block
	mode := cipher.NewCBCEncrypter(block, keyBytes[:aes.BlockSize])

	// Encrypt the input text
	ciphertext := make([]byte, len(inputBytes))
	mode.CryptBlocks(ciphertext, inputBytes)

	// Encode the result in base64
	encryptedText := base64.StdEncoding.EncodeToString(ciphertext)

	return encryptedText
}

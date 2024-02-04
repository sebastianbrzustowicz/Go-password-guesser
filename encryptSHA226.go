package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func encryptSHA224(input string) string {
	// Convert string to bytes
	data := []byte(input)

	// Create a new SHA-224 object
	hash := sha256.New224()

	// Encrypt data
	_, err := hash.Write(data)
	if err != nil {
		return ""
	}

	// Get encrypted data in bytes
	encryptedData := hash.Sum(nil)

	// Convert to hex
	encryptedString := hex.EncodeToString(encryptedData)

	return encryptedString
}

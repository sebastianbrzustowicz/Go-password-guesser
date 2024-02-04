package main

import (
	"crypto/sha512"
	"encoding/hex"
)

func encryptSHA384(input string) string {
	// Convert the string to bytes
	data := []byte(input)

	// Create a new SHA-384 hash object
	hash := sha512.New384()

	// Encrypt the data
	_, err := hash.Write(data)
	if err != nil {
		return ""
	}

	// Get the encrypted data in byte form
	encryptedData := hash.Sum(nil)

	// Convert to hex
	encryptedString := hex.EncodeToString(encryptedData)

	return encryptedString
}

package main

import (
	"encoding/hex"

	"golang.org/x/crypto/md4"
)

func encryptMD4(input string) string {
	// Convert the string to bytes
	data := []byte(input)

	// Create a new MD4 hash object
	hash := md4.New()

	// Encrypt the data
	_, err := hash.Write(data)
	if err != nil {
		return ""
	}

	// Get the encrypted data in byte form
	encryptedData := hash.Sum(nil)

	// Convert the encrypted data to hexadecimal string
	encryptedString := hex.EncodeToString(encryptedData)

	return encryptedString
}

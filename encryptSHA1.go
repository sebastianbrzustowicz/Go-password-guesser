package main

import (
	"crypto/sha1"
	"encoding/hex"
)

func encryptSHA1(input string) string {
	// Create new SHA-1 object
	hasher := sha1.New()

	// Convert string to byte and actualize SHA-1 object
	hasher.Write([]byte(input))

	// Get encrypted result as bytes
	hashBytes := hasher.Sum(nil)

	// Convert to hex
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

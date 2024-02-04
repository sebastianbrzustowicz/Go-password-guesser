package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func encryptSHA256(input string) string {
	// Create new SHA-256 object
	hasher := sha256.New()

	// Convert string to byte and actualize SHA-256 object
	hasher.Write([]byte(input))

	// Get encrypted result as bytes
	hashBytes := hasher.Sum(nil)

	// Convert to hex
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

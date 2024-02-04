package main

import (
	"crypto/md5"
	"encoding/hex"
)

func encryptMD5(input string) string {
	// Create new MD5 object
	hasher := md5.New()

	// Convert string to byte and actualize MD5 object
	hasher.Write([]byte(input))

	// Get encrypted result as bytes
	hashBytes := hasher.Sum(nil)

	// Convert to hex
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

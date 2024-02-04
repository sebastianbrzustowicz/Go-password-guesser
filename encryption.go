// encryption.go

package main

import (
	"fmt"
)

func encryption(target, method string) (encrypted_pw string) {
	switch method {
	case "raw":
		return target
	case "MD5":
		return encryptMD5(target)
	case "SHA-1":
		fmt.Println("SHA-1 method")
		// todo computation
		encryptedPw := "189ewjdwejoer83u38o"
		return encryptedPw
	case "SHA-256":
		fmt.Println("SHA-256 method")
		// todo computation
		encryptedPw := "189ewjdwejoer83u38o"
		return encryptedPw
	case "AES":
		fmt.Println("AES method")
		// todo computation
		encryptedPw := "189ewjdwejoer83u38o"
		return encryptedPw
	case "bcrypt":
		fmt.Println("bcrypt method")
		// todo computation
		encryptedPw := "189ewjdwejoer83u38o"
		return encryptedPw
	case "whirlpool":
		fmt.Println("Whirlpool method")
		// todo computation
		encryptedPw := "189ewjdwejoer83u38o"
		return encryptedPw
	default:
		fmt.Println("Raw password method")
		return target
	}
}

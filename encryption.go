// encryption.go

package main

func encryption(target, method string) (encrypted_pw string) {
	switch method {
	case "raw":
		return target
	case "MD4":
		return encryptMD4(target)
	case "MD5":
		return encryptMD5(target)
	case "SHA-1":
		return encryptSHA1(target)
	case "SHA-224":
		return encryptSHA224(target)
	case "SHA-256":
		return encryptSHA256(target)
	case "SHA-384":
		return encryptSHA384(target)
	case "SHA-512":
		return encryptSHA512(target)

	case "AES":
		return encryptAES(target, "7asAaA65DA156D24EE2A093222d30d4g") // key needed
	case "bcrypt":
		// no implementation bcz of random factor
		return target
	case "whirlpool":
		// no implementation
		return target
	default:
		//fmt.Println("Raw password method")
		return target
	}
}

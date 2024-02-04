package main

import (
	"fmt"
	"os"
	"sync"
)

func worker(targetEncrypted, line, method string, wg *sync.WaitGroup) {
	// Function is encrypting potentially correct password and evaluating it
	defer wg.Done()
	// Encryption
	matchEncrypted := encryption(line, method)

	// Evaluation
	if targetEncrypted == matchEncrypted {
		fmt.Println("Password cracked succesfully:")
		fmt.Println(line)
		os.Exit(0)
	}

}

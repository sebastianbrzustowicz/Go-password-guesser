package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func pwCracker() {

	target := "password"
	method := "raw"
	filePath := "passwords.txt"

	// Read arguments: password, method, path do dictionary file
	result := readArguments()
	if len(result) >= 3 {
		target = result[0]
		method = result[1]
		filePath = result[2]
	}

	// Get encrypted password or "raw method" if password is actually encrypted
	//targetEncrypted := encryption(target, method)
	targetEncrypted := target

	// Workers planning
	if method != "all" {
		// Read file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error during opening file:", err)
			return
		}
		defer file.Close()
		var wg sync.WaitGroup
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Evaluate next potentially correct password
			wg.Add(1)
			go worker(targetEncrypted, scanner.Text(), method, &wg)
		}
		wg.Wait()
		fmt.Println("No matches found")
	} else {
		methods := [8]string{"raw", "MD4", "MD5", "SHA-1", "SHA-224", "SHA-256", "SHA-384", "SHA-512"}
		fmt.Printf("Trying methods: %s", methods[0])
		for i := 1; i < len(methods); i++ {
			fmt.Printf(", %s", methods[i])
		}
		fmt.Printf(" \n")
		var wgMethod sync.WaitGroup
		for i := 0; i < len(methods); i++ {
			//fmt.Printf("Method: %s\n", methods[i])
			wgMethod.Add(1)
			go methodWorker(targetEncrypted, filePath, methods[i], &wgMethod)
		}
		wgMethod.Wait()
		fmt.Println("No matches found")
	}

}

func methodWorker(targetEncrypted, filePath, actualMethod string, wgMethod *sync.WaitGroup) {
	defer wgMethod.Done()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error during opening file:", err)
		return
	}
	defer file.Close()
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Evaluate next potentially correct password
		wg.Add(1)
		go worker(targetEncrypted, scanner.Text(), actualMethod, &wg)

	}
	wg.Wait()
}

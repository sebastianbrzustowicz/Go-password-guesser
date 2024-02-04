// readArguments.go

package main

import (
	"os"
)

func readArguments() []string {
	// Collect all args from CLI
	args := os.Args[1:]

	return args
}

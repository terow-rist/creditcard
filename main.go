package main

import (
	"creditcard/validation"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(0)
	}

	switch args[0] {
	case "validate":
		if len(args) != 1 && args[1] == "--stdin" {
			validation.ValidateNumbers(validation.ProcessingStdin())
		} else {
			validation.ValidateNumbers(args[1:])
		}
	}
}

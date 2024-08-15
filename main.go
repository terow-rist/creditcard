package main

import (
	"fmt"
	"os"

	"creditcard/generation"
	"creditcard/validation"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(0)
	}

	switch args[0] {
	case "validate":
		if len(args) == 1 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
		if args[1] == "--stdin" {
			validation.PrintResults(validation.ProcessingStdin())
		} else {
			validation.PrintResults(args[1:])
		}
	case "generate":
		if len(args) == 1 || len(args) > 2 {
			os.Exit(1)
		}

		asterics_counter := generation.AstericsChecking(args[1])
		for _, digit := range args[1][:len(args[1])-int(asterics_counter)] {
			if digit < '0' || digit > '9' {
				os.Exit(1)
			}
		}

		for _, card := range generation.CreatingAllPossibleCards(asterics_counter, args[1]) {
			fmt.Println(card)
		}

	}
}

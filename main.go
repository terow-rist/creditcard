package main

import (
	"fmt"
	"math/rand"
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
			fmt.Println("Enter the card number for validation.")
			os.Exit(0)
		}
		if args[1] == "--stdin" {
			if len(args) > 2 {
				fmt.Fprintln(os.Stderr, "Too much arguments.")
				os.Exit(1)
			}
			validation.PrintResults(validation.ProcessingStdin())
		} else {
			validation.PrintResults(args[1:])
		}
	case "generate":
		if len(args) == 1 || args[1] == "--pick" && len(args) == 2 {
			fmt.Println("Enter the card number with up to 4 asterics in the end for generation. Example: 440043018030****")
			os.Exit(0)
		}
		if (len(args) > 3 && args[1] == "--pick") || (len(args) > 2 && args[1] != "--pick") {
			fmt.Fprintln(os.Stderr, "Too much arguments.")
			os.Exit(1)
		}

		if args[1] == "--pick" {
			asterics_counter := generation.AstericsChecking(args[2])
			if !generation.IsNumber(args[2], asterics_counter) {
				os.Exit(1)
			}
			all_possible_cards := generation.CreatingAllPossibleCards(asterics_counter, args[2])
			fmt.Println(all_possible_cards[rand.Intn(len(all_possible_cards))])
		} else {
			asterics_counter := generation.AstericsChecking(args[1])
			if !generation.IsNumber(args[1], asterics_counter) {
				os.Exit(1)
			}
			for _, card := range generation.CreatingAllPossibleCards(asterics_counter, args[1]) {
				fmt.Println(card)
			}
		}

	}
}

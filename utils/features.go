package utils

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	ErrTooManyArgs   = "Too many arguments."
	ErrNotEnoughArgs = "Not enough arguments."
	ErrNilArgs       = "No arguments were given."
	ErrIncorrectCmd  = "Incorrect command or order for command."
)

func HandleValidation(args []string) {
	CheckErrNilArgs(args, 1)

	if args[1] == "--stdin" {
		CheckErrTooManyArgs(args, 2)
		PrintResults(ProcessingStdin())
	} else {
		PrintResults(args[1:])
	}
}

func HandleGeneration(args []string) {
	CheckErrNilArgs(args, 1)

	if args[1] == "--pick" {
		CheckErrNilArgs(args, 2)
		CheckErrTooManyArgs(args, 3)
		if !IsNumber(args[2]) {
			os.Exit(1)
		}
		all_possible_cards := CreatingAllPossibleCards(args[2])
		fmt.Println(all_possible_cards[rand.Intn(len(all_possible_cards))])
	} else {
		CheckErrTooManyArgs(args, 2)
		if !IsNumber(args[1]) {
			os.Exit(1)
		}
		for _, card := range CreatingAllPossibleCards(args[1]) {
			fmt.Println(card)
		}
	}
}

func HandleInformation(args []string) {
	CheckErrNilArgs(args, 1)
	CheckErrNotEnoughArgs(args, 4)
	CheckErrIncorrectCmd(args[1], "--brands=brands.txt")
	CheckErrIncorrectCmd(args[2], "--issuers=issuers.txt")
	till_the_end := 0
	for _, card := range args[3:] {
		if ValidationConditions(card) {
			fmt.Println(card)
			fmt.Println("Correct: yes")
			fmt.Println("Card Brand:", BrandsCheck(card))
			fmt.Println("Card Issuer:", IssuerCheck(card))
		} else {
			fmt.Println(card)
			fmt.Println("Correct: no")
			fmt.Println("Card Brand: -")
			fmt.Println("Card Issuer: -")
		}
		if till_the_end != len(args[3:])-1 {
			fmt.Println()
		}
		till_the_end++
	}
}

func CheckErrNotEnoughArgs(args []string, max_size int) {
	if len(args) < max_size {
		fmt.Fprintln(os.Stderr, ErrNotEnoughArgs)
		os.Exit(1)
	}
}

func CheckErrTooManyArgs(args []string, max_size int) {
	if len(args) > max_size {
		fmt.Fprintln(os.Stderr, ErrTooManyArgs)
		os.Exit(1)
	}
}

func CheckErrNilArgs(args []string, max_size int) {
	if len(args) == max_size {
		fmt.Fprintln(os.Stderr, ErrNilArgs)
		os.Exit(1)
	}
}

func CheckErrIncorrectCmd(cmd string, excepted_cmd string) {
	if cmd != excepted_cmd {
		fmt.Fprintln(os.Stderr, ErrIncorrectCmd)
		os.Exit(1)
	}
}

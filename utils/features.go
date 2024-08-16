package utils

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	ErrTooManyArgs = "Too many arguments."
	ErrNilArgs     = "No arguments were given."
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

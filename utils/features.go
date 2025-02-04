package utils

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
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
		if len(all_possible_cards) == 0 {
			os.Exit(1)
		}
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

	var cards []string
	if args[3] == "--stdin" {
		cards = ProcessingStdin()
		fmt.Println()
	} else {
		cards = args[3:]
	}
	till_the_end := 0
	for _, card := range cards {
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
		if till_the_end != len(cards)-1 {
			fmt.Println()
		}
		till_the_end++
	}
}

func HandleIssue(args []string) {
	CheckErrNilArgs(args, 1)
	CheckErrNotEnoughArgs(args, 5)
	CheckErrIncorrectCmd(args[1], "--brands=brands.txt")
	CheckErrIncorrectCmd(args[2], "--issuers=issuers.txt")
	if strings.HasPrefix(args[3], "--brand=") && strings.HasPrefix(args[4], "--issuer=") {
		brand, issuer := args[3][8:], args[4][9:]
		brand_num, issuer_num := BrandsNumber(brand), IssuerNumber(issuer)
		if len(brand_num) != 0 && len(issuer_num) != 0 && brand_num[0] == issuer_num[0] {
			fmt.Println(IssueGenerate(issuer_num[0], issuer_num))
		} else {
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, ErrIncorrectCmd)
		os.Exit(1)
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

func ProcessingStdin() []string {
	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	return strings.Fields(input)
}

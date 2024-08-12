package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(0)
	}

	switch args[0] {
	case "validate":
		if args[1] == "--stdin" {
			validateNumbers(processingStdin())
		} else {
			validateNumbers(args[1:])
		}
	}
}

func processingStdin() []string {
	reader := os.Stdin
	buf := new(strings.Builder)
	io.Copy(buf, reader)
	input := buf.String()
	return strings.Fields(input)
}

func validateNumbers(numbers []string) {
	for _, card_num := range numbers {
		if len(card_num) < 13 || len(card_num) > 16 {
			fmt.Println("INCORRECT")
			os.Exit(1)
		}
		for _, digit := range card_num {
			if digit < '0' || digit > '9' {
				fmt.Println("INCORRECT")
				os.Exit(1)
			}
		}
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("OK")
	}
	os.Exit(0)
}

package main

import (
	"fmt"
	"os"
	"strconv"

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
		if len(args) == 1 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
		asterics_counter := 0
		asterics_indexes := []int{}
		for i := len(args[1]) - 1; i >= 0; i-- {
			if args[1][i] == '*' {
				if len(args[1])-5 > i {
					os.Exit(1)
				}
				asterics_counter++
				asterics_indexes = append(asterics_indexes, i)
			}
			if asterics_counter > 4 {
				os.Exit(1)
			}
		}
		var card_with_asterics string
		switch asterics_counter {
		case 0:
			card_with_asterics = args[1][:len(args[1])]
		case 1:
			card_with_asterics = args[1][:len(args[1])-1] + "0"
		case 2:
			card_with_asterics = args[1][:len(args[1])-2] + "00"
		case 3:
			card_with_asterics = args[1][:len(args[1])-3] + "000"
		default:
			card_with_asterics = args[1][:len(args[1])-4] + "0000"
		}

		possible_card, _ := strconv.ParseInt(card_with_asterics, 10, 64)
		fmt.Println(possible_card)

	}
}

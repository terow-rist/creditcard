package main

import (
	"fmt"
	"math"
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
		asterics_counter := 0.0

		end_of_asterics := false
		for i := len(args[1]) - 1; i >= 0; i-- {
			if args[1][i] == '*' {
				if end_of_asterics {
					os.Exit(1)
				}
				asterics_counter++
			} else {
				end_of_asterics = true
			}
			if asterics_counter > 4 {
				os.Exit(1)
			}
		}
		var card_from_stdin string
		switch asterics_counter {
		case 0:
			card_from_stdin = args[1][:len(args[1])]
		case 1:
			card_from_stdin = args[1][:len(args[1])-1] + "0"
		case 2:
			card_from_stdin = args[1][:len(args[1])-2] + "00"
		case 3:
			card_from_stdin = args[1][:len(args[1])-3] + "000"
		default:
			card_from_stdin = args[1][:len(args[1])-4] + "0000"
		}

		card_in_int64, _ := strconv.ParseInt(card_from_stdin, 10, 64)
		possible_cards := []string{}
		for i := 0; i < int(math.Pow(10, asterics_counter)); i++ {
			if validation.ValidationConditions(card_from_stdin) {
				possible_cards = append(possible_cards, card_from_stdin)
			}
			card_in_int64++
			card_from_stdin = strconv.Itoa(int(card_in_int64))
		}

		for _, card := range possible_cards {
			fmt.Println(card)
		}
	}
}

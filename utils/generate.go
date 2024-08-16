package utils

import (
	"math"
	"os"
	"strconv"
)

func AstericsChecking(input string) float64 {
	asterics_counter := 0.0

	end_of_asterics := false
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '*' {
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
	return asterics_counter
}

func CreatingAllPossibleCards(input string) []string {
	asterics_counter := AstericsChecking(input)
	var card_from_stdin string
	switch asterics_counter {
	case 0:
		card_from_stdin = input[:]
	case 1:
		card_from_stdin = input[:len(input)-1] + "0"
	case 2:
		card_from_stdin = input[:len(input)-2] + "00"
	case 3:
		card_from_stdin = input[:len(input)-3] + "000"
	default:
		card_from_stdin = input[:len(input)-4] + "0000"
	}

	card_in_int64, _ := strconv.ParseInt(card_from_stdin, 10, 64)
	possible_cards := []string{}
	for i := 0; i < int(math.Pow(10, asterics_counter)); i++ {
		if ValidationConditions(card_from_stdin) {
			possible_cards = append(possible_cards, card_from_stdin)
		}
		card_in_int64++
		card_from_stdin = strconv.Itoa(int(card_in_int64))
	}
	return possible_cards
}

func IsNumber(input string) bool {
	asterics_counter := AstericsChecking(input)
	for _, digit := range input[:len(input)-int(asterics_counter)] {
		if digit < '0' || digit > '9' {
			return false
		}
	}
	return true
}

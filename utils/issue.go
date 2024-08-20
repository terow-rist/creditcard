package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func BrandsNumber(brand string) string {
	content, err := os.ReadFile("brands.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading file.")
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		slc_of_brand := strings.Split(line, ":")
		if slc_of_brand[0] == brand {
			return slc_of_brand[1]
		}
	}
	return ""
}

func IssuerNumber(issuer string) string {
	content, err := os.ReadFile("issuers.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading file.")
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		slc_of_brand := strings.Split(line, ":")
		if slc_of_brand[0] == issuer {
			return slc_of_brand[1]
		}
	}
	return ""
}

func IssueGenerate(first_digit byte, card_start string) string {
	if first_digit == '4' {
		slc_with_card_len := []int{7, 10}
		for {
			card_end := card_start
			for i := 0; i < slc_with_card_len[rand.Intn(2)]; i++ {
				card_end += string(rune('0' + rand.Intn(10)))
			}
			if ValidationConditions(card_end) {
				return card_end
			}
		}
	} else if first_digit == '5' {
		for {
			card_end := card_start
			for i := 0; i < 10; i++ {
				card_end += string(rune('0' + rand.Intn(10)))
			}
			if ValidationConditions(card_end) {
				return card_end
			}
		}
	} else if first_digit == '3' {
		card_end := card_start
		for i := 0; i < 9; i++ {
			card_end += string(rune('0' + rand.Intn(10)))
		}
		if ValidationConditions(card_end) {
			return card_end
		}
	}
	return ""
}

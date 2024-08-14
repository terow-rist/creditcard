package validation

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcessingStdin() []string {
	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	return strings.Fields(input)
}

func ValidateNumbers(numbers []string) int {
	for index, card_num := range numbers {
		if !ValidationConditions(card_num) {
			return index
		} else if !LuhnAlgorithm(card_num) {
			return index
		}
		for _, digit := range card_num {
			if digit < '0' || digit > '9' {
				return index
			}
		}
	}
	return len(numbers)
}

func LuhnAlgorithm(str string) bool {
	var sum int
	is2 := false
	for i := len(str) - 1; i >= 0; i-- {
		if is2 {
			if int(str[i]-'0')*2 > 9 {
				sum += int(str[i]-'0')*2 - 9
			} else {
				sum += int(str[i]-'0') * 2
			}
		} else {
			sum += int(str[i] - '0')
		}
		is2 = !is2
	}
	return sum%10 == 0
}

func PrintResults(numbers []string) {
	for i := range numbers {
		if ValidateNumbers(numbers) == i {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
		fmt.Println("OK")
	}
	os.Exit(0)
}

func ValidationConditions(card_num string) bool {
	if card_num[0] < '3' || card_num[0] > '5' {
		return false
	} else if len(card_num) < 13 || len(card_num) > 16 {
		return false
	}

	if card_num[0] == '4' && len(card_num) != 13 && len(card_num) != 16 {
		return false
	}
	if card_num[0] == '5' {
		if card_num[1] == '1' || card_num[1] == '2' || card_num[1] == '3' || card_num[1] == '4' || card_num[1] == '5' {
			if len(card_num) != 16 {
				return false
			}
		} else {
			return false
		}
	}
	if card_num[0] == '3' {
		if card_num[1] == '7' || card_num[1] == '4' {
			if len(card_num) != 15 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

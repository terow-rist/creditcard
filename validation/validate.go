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

func ValidateNumbers(numbers []string) {
	for _, card_num := range numbers {
		if card_num[0] < '3' || card_num[0] > '5' {
			ErrorOutput()
		} else if len(card_num) < 13 || len(card_num) > 16 {
			ErrorOutput()
		} else if !LuhnAlgorithm(card_num) {
			ErrorOutput()
		}

		for _, digit := range card_num {
			if digit < '0' || digit > '9' {
				ErrorOutput()
			}
		}
		if card_num[0] == '4' && len(card_num) != 13 && len(card_num) != 16 {
			ErrorOutput()
		}
		if card_num[0] == '5' {
			if card_num[1] == '1' || card_num[1] == '2' || card_num[1] == '3' || card_num[1] == '4' || card_num[1] == '5' {
				if len(card_num) != 16 {
					ErrorOutput()
				}
			} else {
				ErrorOutput()
			}
		}
		if card_num[0] == '3' {
			if card_num[1] == '7' || card_num[1] == '4' {
				if len(card_num) != 15 {
					ErrorOutput()
				}
			} else {
				ErrorOutput()
			}
		}
		fmt.Println("OK")
	}
	os.Exit(0)
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

func ErrorOutput() {
	fmt.Fprintln(os.Stderr, "INCORRECT")
	os.Exit(1)
}

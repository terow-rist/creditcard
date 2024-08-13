package validation

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcessingStdin() []string {
	reader := os.Stdin
	buf := new(strings.Builder)
	io.Copy(buf, reader)
	input := buf.String()
	return strings.Fields(input)
}

func ValidateNumbers(numbers []string) {
	for _, card_num := range numbers {
		if !LuhnAlgorithm(card_num) {
			fmt.Println("INCORRECT")
			os.Exit(1)
		}
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
	if sum%10 == 0 {
		return true
	}
	return false
}

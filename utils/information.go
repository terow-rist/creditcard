package utils

import (
	"fmt"
	"os"
	"strings"
)

func BrandsCheck(brand string, card string) string {
	if brand != "--brands=brands.txt" {
		fmt.Fprintln(os.Stderr, "Inccorect command.")
		os.Exit(1)
	}

	content, err := os.ReadFile("brands.txt")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading file.")
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	brands := map[string]string{}
	till_the_end := 0

	for _, line := range lines {
		if till_the_end > len(lines)-1 {
			break
		}
		slc_of_brand := strings.Split(line, ":")
		brands[string(slc_of_brand[1][0])] = slc_of_brand[0]
		till_the_end++
	}

	if brands[string(card[0])] == "" {
		return "-"
	}
	return brands[string(card[0])]

}

package utils

import (
	"fmt"
	"os"
	"strings"
)

func BrandsCheck(card string) string {
	content, err := os.ReadFile("brands.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading file.")
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	brands := map[string]string{}

	for _, line := range lines {
		slc_of_brand := strings.Split(line, ":")
		brands[string(slc_of_brand[1][0])] = slc_of_brand[0]
	}

	if brands[string(card[0])] == "" {
		return "-"
	}
	return brands[string(card[0])]
}

func IssuerCheck(card string) string {
	content, err := os.ReadFile("issuers.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading file.")
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		slc_of_issuer := strings.Split(line, ":")
		if strings.HasPrefix(card, slc_of_issuer[1]) {
			return slc_of_issuer[0]
		}
	}
	return "-"
}

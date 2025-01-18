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
		if len(slc_of_brand) < 2 {
			continue
		}
		if len(slc_of_brand[1]) > 0 {
			brands[string(slc_of_brand[1][0])] = slc_of_brand[0]
		}
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
		if len(slc_of_issuer) < 2 {
			continue
		}
		if strings.HasPrefix(card, slc_of_issuer[1]) {
			return slc_of_issuer[0]
		}
	}
	return "-"
}

/*
Examples:
	4400430180300003 4042430180300007 5177920000000005 4405630000005 awd 5395450000000009
*/

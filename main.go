package main

import (
	"os"

	"creditcard/utils"
)

func main() {
	args := os.Args[1:]

	utils.CheckErrNilArgs(args, 0)

	switch args[0] {
	case "validate":
		utils.HandleValidation(args)
	case "generate":
		utils.HandleGeneration(args)
	case "information":
		utils.HandleInformation(args)
	case "issue":
		utils.HandleIssue(args)
	default:
		os.Exit(1)
	}
}

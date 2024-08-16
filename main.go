package main

import (
	"creditcard/utils"
	"fmt"
	"os"
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
		fmt.Println(utils.BrandsCheck(args[1], args[2]))
	}

}

package cmd

import (
	"fmt"
)

func Report(args []string) {
	argsLen := len(args)
	reportType := ""

	if argsLen == 0 {
		reportType = "list"
	} else {
		reportType = args[0]
	}
	fmt.Println("reportType", reportType)
}

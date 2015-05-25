package main

import (
	"github.com/jlgerber/penv/pack"
	"os"
)

func main() {
	args := os.Args
	searchterm := ""
	if len(args) >= 2 {
		searchterm = args[1]
	}

	penv.PrintEnv(searchterm, penv.ReplaceColon, penv.FormatPrint)

}

package main

import (
	"github.com/codegangsta/cli"
	"github.com/jlgerber/penv/pack"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "penv"
	app.Usage = penv.Usage
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{

		cli.BoolFlag{
			Name:  "pretty, p",
			Usage: "pretty print the result. Any colon separated paths will be expanded to print a path per line.",
		},
		cli.BoolFlag{
			Name:  "exact, e",
			Usage: "Perform an exact match when provided with a matching argument.",
		},
	}

	app.Action = func(c *cli.Context) {
		args := c.Args()
		searchterm := ""

		if len(args) >= 1 {
			searchterm = args[0]
		}
		// evaluate options
		printfunc := penv.PassThroughVMF
		formatfunc := penv.FormatPrint
		matchfunc := penv.ContainsMatch

		if c.Bool("pretty") == true {
			printfunc = penv.ReplaceColon
		}

		if c.Bool("exact") == true {
			matchfunc = penv.ExactMatch
		}

		penv.PrintEnv(searchterm, matchfunc, printfunc, formatfunc)

	}

	app.Run(os.Args)

}

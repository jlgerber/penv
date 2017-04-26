package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/jlgerber/penv/pack"
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
		cli.BoolFlag{
			Name:  "sort, s",
			Usage: "Sort the resulting keys. Note that the paths will NOT be sorted when using the --pretty flag; Only the environment variable names are sorted in the event that multiple variables match the provided inputs",
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

		penv.PrintEnv(searchterm, matchfunc, printfunc, formatfunc, c.Bool("sort"))

	}

	app.Run(os.Args)

}

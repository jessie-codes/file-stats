package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "file-stats"

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "input, i",
			Usage: "`FILE` to get statistics for",
		},
		cli.StringFlag{
			Name:  "keyword, k",
			Usage: "Load keywords from `FILE`",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "`FILE` to save statistics data",
		},
	}

	app.Action = action

	app.Run(os.Args)

}

func action(c *cli.Context) {
	loadKeywords(c.String("keyword"))
	readFiles(c.StringSlice("input"))
	logResults(c.String("output"))
}

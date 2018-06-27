package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "file-stats"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "`GLOB_PATTERN` of files to get statistics for",
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
	newStatistics().loadKeywords(c.String("keyword")).readFiles(c.String("input")).calc().output(c.String("output"))
}

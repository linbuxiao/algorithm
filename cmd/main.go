package main

import (
	"github.com/linbuxiao/algorithm/cmd/action"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "Algorithm",
		Commands: []*cli.Command{
			action.NewProblem,
			action.Get,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

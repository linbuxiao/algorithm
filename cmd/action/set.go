package action

import (
	"errors"
	"fmt"
	"github.com/linbuxiao/algorithm/pkg/app"
	"github.com/linbuxiao/algorithm/pkg/model"
	"github.com/linbuxiao/algorithm/pkg/util"
	"github.com/urfave/cli/v2"
)

var Set = &cli.Command{
	Name: "set",
	Subcommands: []*cli.Command{
		{
			Name: "status",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "problem",
					Aliases: []string{"p"},
				},
				&cli.StringFlag{
					Name:    "namespace",
					Aliases: []string{"ns"},
				},
			},
			Action: func(ctx *cli.Context) error {
				p := ctx.String("problem")
				ns := ctx.String("namespace")
				statusStr := ctx.Args().First()
				status := app.UnFinished
				if statusStr == "true" {
					status = app.Finished
				} else if statusStr != "false" {
					return errors.New("wrong status")
				}
				if ns == "" {
					ns = app.DefaultNameSpace
				}
				repo, err := model.GetProblemRepo(util.GetLeetcodeUrlFromFileName(p), ns)
				if err != nil {
					return err
				}
				fmt.Println(status)
				if err := repo.SetStatus(status); err != nil {
					return err
				}
				return nil
			},
		},
	},
}

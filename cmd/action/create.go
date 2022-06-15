package action

import (
	"github.com/linbuxiao/algorithm/pkg/app"
	"github.com/linbuxiao/algorithm/pkg/model"
	"github.com/urfave/cli/v2"
)

var Create = &cli.Command{
	Name: "create",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "namespace",
			Aliases: []string{"ns"},
		},
	},
	Action: func(ctx *cli.Context) error {
		url := ctx.Args().First()
		ns := ctx.String("namespace")
		if ns == "" {
			ns = app.DefaultNameSpace
		}
		problem := model.NewProblemRepo(url, ns)
		if err := problem.CreateProblemFile(); err != nil {
			return err
		}
		return nil
	},
}

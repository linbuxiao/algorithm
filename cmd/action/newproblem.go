package action

import (
	"github.com/linbuxiao/algorithm/pkg/model"
	"github.com/urfave/cli/v2"
)

const defaultNameSpace = "default"

var NewProblem = &cli.Command{
	Name: "newproblem",
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
			ns = defaultNameSpace
		}
		problem := model.NewProblemRepo(url, ns)
		if err := problem.CreateProblemFile(); err != nil {
			return err
		}
		return nil
	},
}

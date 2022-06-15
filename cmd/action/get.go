package action

import (
	"fmt"
	"github.com/linbuxiao/algorithm/pkg/util"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"os"
)

var Get = &cli.Command{
	Name: "get",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "namespace",
			Aliases: []string{"ns"},
		},
		&cli.StringFlag{
			Name:    "problem",
			Aliases: []string{"p"},
		},
	},
	Action: func(ctx *cli.Context) error {
		setNS := ctx.String("namespace")
		setProblem := ctx.String("problem")
		nameSpaceAbsPathArr, err := util.GetAllNameSpacePath()
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"problem",
			"finished",
			"name space",
			"url",
		})
		for _, v := range nameSpaceAbsPathArr {
			ns := util.GetNameByAbsPath(v)
			if setNS != "" && ns != setNS {
				continue
			}
			problemAbsPathArr, err := util.GetProblemsByNameSpace(ns)
			if err != nil {
				return err
			}
			for _, k := range problemAbsPathArr {
				problemName, status := util.GetProblemFromFileName(util.GetNameByAbsPath(k))
				statusStr := "false"
				if status {
					statusStr = "true"
				}
				if setProblem != "" && problemName != setProblem {
					continue
				}
				table.Append([]string{
					problemName,
					statusStr,
					ns,
					fmt.Sprintf("https://leetcode.cn/problems/%s/", util.TransformFileNameToSlug(problemName)),
				})
			}
		}
		table.Render()
		return nil
	},
}

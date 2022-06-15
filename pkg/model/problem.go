package model

import (
	"bufio"
	"fmt"
	"github.com/linbuxiao/algorithm/pkg/app"
	"github.com/linbuxiao/algorithm/pkg/util"
	"html/template"
	"io/ioutil"
	"os"
)

type ProblemTmplRepo struct {
	Url             string
	NameSpaceDir    string
	ProblemFileName string
	NameSpace       string
}

func NewProblemRepo(url string, namespace string) *ProblemTmplRepo {
	problemDir := fmt.Sprintf("%s%s", util.GetRootPath(), app.ProblemDir)
	nameSpaceDir := fmt.Sprintf("%s/%s", problemDir, namespace)
	slug := util.GetSlugFromLeetcodeUrl(url)
	problemFileName := fmt.Sprintf("%s/%s", nameSpaceDir, util.SetProblemFileName(util.TransformSlugToFileName(slug), false))
	p := &ProblemTmplRepo{
		Url:             url,
		NameSpace:       namespace,
		NameSpaceDir:    nameSpaceDir,
		ProblemFileName: problemFileName,
	}
	return p
}

func (p *ProblemTmplRepo) CreateProblemFile() error {
	if err := p.makeSureNameSpaceExist(); err != nil {
		return err
	}
	file, err := os.Create(p.ProblemFileName)
	if err != nil {
		return fmt.Errorf("create %s failed: %+v", p.ProblemFileName, err)
	}
	w := bufio.NewWriter(file)
	tmplPath := fmt.Sprintf("%s%s", util.GetRootPath(), app.ProblemTemplatePath)
	tmplFile, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		return err
	}
	if err = template.Must(template.New(app.ProblemTemplateName).Parse(string(tmplFile))).Execute(w, p); err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	return nil
}

func (p *ProblemTmplRepo) makeSureNameSpaceExist() error {
	if !util.DirExist(p.NameSpaceDir) {
		if err := os.Mkdir(p.NameSpaceDir, 0755); err != nil {
			return fmt.Errorf("mkdir '%s' for writing failed: %+v\n", p.NameSpaceDir, err)
		}
	}
	return nil
}

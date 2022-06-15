package model

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/linbuxiao/algorithm/pkg/app"
	"github.com/linbuxiao/algorithm/pkg/util"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type ProblemTmplRepo struct {
	Url             string
	NameSpaceDir    string
	ProblemFileName string
	NameSpace       string
	ProblemName     string
	Status          bool
	Slug            string
}

func NewProblemRepo(url string, namespace string) *ProblemTmplRepo {
	p := getProblemRepo(url, namespace)
	problemFileName := fmt.Sprintf("%s/%s", p.NameSpaceDir, util.GetProblemFileName(p.ProblemName, app.UnFinished))
	p.Status = false
	p.ProblemFileName = problemFileName
	return p
}

func GetProblemRepo(url string, namespace string) (*ProblemTmplRepo, error) {
	p := getProblemRepo(url, namespace)
	ok, err := p.isExist()
	if err != nil {
		return nil, fmt.Errorf("cannot get problem exist: %+v", err)
	}
	if !ok {
		return nil, errors.New("there is no such problem")
	}
	return p, nil
}

func getProblemRepo(url string, namespace string) *ProblemTmplRepo {
	problemDir := fmt.Sprintf("%s%s", util.GetRootPath(), app.ProblemDir)
	nameSpaceDir := fmt.Sprintf("%s/%s", problemDir, namespace)
	slug := util.GetSlugFromLeetcodeUrl(url)
	problemName := util.TransformSlugToFileName(slug)
	return &ProblemTmplRepo{
		Url:          url,
		NameSpace:    namespace,
		NameSpaceDir: nameSpaceDir,
		ProblemName:  problemName,
		Slug:         slug,
	}
}

func (p *ProblemTmplRepo) CreateProblemFile() error {
	isExist, err := p.isExist()
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("repeat question in same namespace")
	}
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

func (p *ProblemTmplRepo) isExist() (bool, error) {
	files, err := util.GetProblemsByNameSpace(p.NameSpace)
	if err != nil {
		return false, err
	}
	for _, v := range files {
		if strings.Contains(v, fmt.Sprintf("%s.", p.ProblemName)) {
			if err := p.syncStatus(v); err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (p *ProblemTmplRepo) SetStatus(status bool) error {
	if err := os.Rename(p.ProblemFileName, fmt.Sprintf("%s/%s", p.NameSpaceDir, util.GetProblemFileName(p.ProblemName, status))); err != nil {
		return err
	}
	p.Status = status
	return nil
}

func (p *ProblemTmplRepo) syncStatus(targetProblem string) error {
	p.ProblemFileName = targetProblem
	var err error
	if util.GetProblemFileName(p.ProblemName, app.UnFinished) == targetProblem {
		err = p.SetStatus(app.UnFinished)
	} else {
		err = p.SetStatus(app.Finished)
	}
	return err
}

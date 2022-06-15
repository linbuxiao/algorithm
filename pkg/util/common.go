package util

import (
	"fmt"
	"github.com/linbuxiao/algorithm/pkg/app"
	"os"
	"path/filepath"
	"strings"
)

func GetSlugFromLeetcodeUrl(url string) string {
	arr := strings.Split(url, "/")
	return arr[len(arr)-2]
}

func TransformSlugToFileName(s string) string {
	return strings.ReplaceAll(s, "-", "_")
}

func TransformFileNameToSlug(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

func DirExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GetRootPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	index := strings.Index(wd, app.PackageName)
	find := false
	var k int
	for i := index; i <= len([]rune(wd))-1; i++ {
		if wd[i] == os.PathSeparator {
			k = i
			find = true
			break
		}
	}
	if find {
		return wd[:k]
	} else {
		return wd
	}
}

func GetAllNameSpacePath() ([]string, error) {
	return filepath.Glob(fmt.Sprintf("%s%s/*", GetRootPath(), app.ProblemDir))
}

func GetProblemsByNameSpace(ns string) ([]string, error) {
	nsAbsPath := fmt.Sprintf("%s%s/%s", GetRootPath(), app.ProblemDir, ns)
	return filepath.Glob(fmt.Sprintf("%s/*", nsAbsPath))
}

func GetNameByAbsPath(abs string) string {
	arr := strings.Split(abs, "/")
	return arr[len(arr)-1]
}

func SetProblemFileName(name string, finished bool) string {
	if finished {
		return fmt.Sprintf("%s.true.go", name)
	}
	return fmt.Sprintf("%s.false.go", name)
}

func GetProblemFromFileName(name string) (string, bool) {
	if strings.HasSuffix(name, app.FinishedProblemSuffix) {
		return strings.TrimSuffix(name, app.FinishedProblemSuffix), true
	}
	return strings.TrimSuffix(name, app.UnFinishedProblemSuffix), true
}

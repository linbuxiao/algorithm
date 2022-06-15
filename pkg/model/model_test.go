package model_test

import (
	"github.com/linbuxiao/algorithm/pkg/model"
	"github.com/linbuxiao/algorithm/pkg/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Model", func() {
	Describe("new problem instance", func() {
		Context("make sure new instance", func() {
			const url = "https://leetcode.cn/problems/test-slug/"
			const ns = "test-namespace"
			It("should return new instance by rules", func() {
				p := model.NewProblemRepo(url, ns)
				Expect(p.Url).To(Equal(url))
				Expect(p.NameSpace).To(Equal(ns))
			})
		})

		Context("no namespace", func() {
			const (
				url              = "https://leetcode.cn/problems/test-slug/"
				testNameSpaceDir = "./test-namespace"
				testProblemFile  = "./test-namespace/test_slug.go"
				testNameSpace    = "test-namespace"
			)

			var p *model.ProblemTmplRepo
			BeforeEach(func() {
				p = &model.ProblemTmplRepo{
					Url:             url,
					NameSpace:       testNameSpace,
					NameSpaceDir:    testNameSpaceDir,
					ProblemFileName: testProblemFile,
				}
			})
			AfterEach(func() {
				err := os.RemoveAll(testNameSpaceDir)
				Expect(err).To(BeNil())
			})
			It("should create a namespace", func() {
				err := p.CreateProblemFile()
				Expect(err).To(BeNil())
				Expect(util.DirExist(testNameSpaceDir)).To(BeTrue())
			})
		})
	})
})

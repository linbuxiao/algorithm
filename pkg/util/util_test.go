package util_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"strings"

	"github.com/linbuxiao/algorithm/pkg/util"
)

var _ = Describe("Util", func() {
	Describe("get slug from url", func() {
		Context("url: https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/", func() {
			It("should be fei-bo-na-qi-shu-lie-lcof", func() {
				Expect(util.GetSlugFromLeetcodeUrl("https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/")).To(Equal("fei-bo-na-qi-shu-lie-lcof"))
			})
		})
	})

	Describe("make sure dir", func() {
		Context("no exsit dir", func() {
			It("should return false because of there is no dir", func() {
				Expect(util.DirExist("./no-exist-dir")).To(BeFalse())
			})
		})

		Context("create dir", func() {
			BeforeEach(func() {
				err := os.MkdirAll("./test-dir", os.ModePerm)
				Expect(err).To(BeNil())
			})
			AfterEach(func() {
				err := os.RemoveAll("./test-dir")
				Expect(err).To(BeNil())
			})
			It("should return true because of there has dir", func() {
				Expect(util.DirExist("./test-dir")).To(BeTrue())
			})
		})
	})

	Describe("get wd", func() {
		It("should return a string path with algorithm", func() {
			Expect(strings.HasSuffix(util.GetRootPath(), "algorithm")).To(BeTrue())
		})
	})
})

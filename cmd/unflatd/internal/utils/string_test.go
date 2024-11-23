package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/utils"
)

var _ = Describe("String Utils", func() {
	Describe("Contains", func() {
		It("returns true when slice contains any item from items", func() {
			slice := []string{"apple", "banana", "orange"}
			items := []string{"banana", "grape"}
			Expect(utils.Contains(slice, items)).To(BeTrue())
		})

		It("returns false when slice contains none of the items", func() {
			slice := []string{"apple", "banana", "orange"}
			items := []string{"grape", "kiwi"}
			Expect(utils.Contains(slice, items)).To(BeFalse())
		})

		It("returns false with empty items slice", func() {
			slice := []string{"apple", "banana", "orange"}
			items := []string{}
			Expect(utils.Contains(slice, items)).To(BeFalse())
		})

		It("returns false with empty source slice", func() {
			slice := []string{}
			items := []string{"apple", "banana"}
			Expect(utils.Contains(slice, items)).To(BeFalse())
		})
	})
})

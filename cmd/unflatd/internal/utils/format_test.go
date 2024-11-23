package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/utils"
)

var _ = Describe("Format Utils", func() {
	Describe("ToSnakeCase", func() {
		DescribeTable("converts strings to snake case",
			func(input, expected string) {
				Expect(utils.ToSnakeCase(input)).To(Equal(expected))
			},
			Entry("camelCase", "camelCase", "camel_case"),
			Entry("PascalCase", "PascalCase", "pascal_case"),
			Entry("mixed_Case", "mixedCAPS", "mixed_caps"),
			Entry("multiple words", "ThisIsATest", "this_is_a_test"),
			Entry("single word", "word", "word"),
			Entry("acronyms", "APIResponse", "api_response"),
			Entry("numbers", "User123Name", "user123_name"),
		)
	})

	Describe("Zip", func() {
		Context("with string slices", func() {
			It("creates a map from two equal-length slices", func() {
				keys := []string{"a", "b", "c"}
				values := []string{"1", "2", "3"}
				result := utils.Zip(keys, values)

				expected := map[string]string{
					"a": "1",
					"b": "2",
					"c": "3",
				}
				Expect(result).To(Equal(expected))
			})
		})

		Context("with integer slices", func() {
			It("creates a map from two equal-length slices", func() {
				keys := []int{1, 2, 3}
				values := []int{10, 20, 30}
				result := utils.Zip(keys, values)

				expected := map[int]int{
					1: 10,
					2: 20,
					3: 30,
				}
				Expect(result).To(Equal(expected))
			})
		})

		Context("with unequal length slices", func() {
			It("panics when slice lengths don't match", func() {
				keys := []string{"a", "b"}
				values := []string{"1"}

				Expect(func() {
					utils.Zip(keys, values)
				}).To(PanicWith("slice lengths do not match: 2 != 1"))
			})
		})

		Context("with empty slices", func() {
			It("returns an empty map for empty slices", func() {
				keys := []string{}
				values := []string{}
				result := utils.Zip(keys, values)
				Expect(result).To(BeEmpty())
			})
		})
	})
})

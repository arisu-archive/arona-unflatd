package conversion_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/conversion"
	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
)

var _ = Describe("FieldType", func() {
	Context("ConvertFieldType", func() {
		It("should handle nullable types", func() {
			result := conversion.ConvertFieldType("Nullable<int>")
			Expect(result).To(Equal("int"))
		})

		It("should pass through regular types", func() {
			result := conversion.ConvertFieldType("string")
			Expect(result).To(Equal("string"))
		})
	})

	Context("ConvertEnumValues", func() {
		It("should convert int enum values", func() {
			values := map[string]string{
				"None":   "0",
				"First":  "1",
				"Second": "2",
			}
			result := conversion.ConvertEnumValues("int", values)
			Expect(result).To(Equal([]fbs.EnumValue{
				{Name: "None", Value: int64(0)},
				{Name: "First", Value: int64(1)},
				{Name: "Second", Value: int64(2)},
			}))
		})

		It("should convert uint enum values", func() {
			values := map[string]string{
				"None":   "0",
				"First":  "1",
				"Second": "2",
			}
			result := conversion.ConvertEnumValues("uint", values)
			Expect(result).To(Equal([]fbs.EnumValue{
				{Name: "None", Value: uint64(0)},
				{Name: "First", Value: uint64(1)},
				{Name: "Second", Value: uint64(2)},
			}))
		})
	})
})

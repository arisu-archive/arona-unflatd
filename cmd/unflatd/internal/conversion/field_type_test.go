package conversion_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/conversion"
	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
)

var _ = Describe("FieldType", func() {
	Context("ExtractNamespaceType", func() {
		It("should extract namespace and name correctly", func() {
			result := conversion.ExtractNamespace("MyNamespace.MyType")
			Expect(result).To(Equal("MyNamespace"))
		})

		It("should handle types without namespace", func() {
			result := conversion.ExtractNamespace("MyType")
			Expect(result).To(Equal(""))
		})

		It("should handle empty string", func() {
			result := conversion.ExtractNamespace("")
			Expect(result).To(Equal(""))
		})

		It("should handle multiple dots in type", func() {
			result := conversion.ExtractNamespace("A.B.C.MyType")
			Expect(result).To(Equal("A.B.C"))
		})
	})

	Context("ConvertFieldType", func() {
		It("should handle nullable types", func() {
			result := conversion.ConvertFieldType("Nullable<int>")
			Expect(result).To(Equal("int"))
		})

		It("should pass through regular types", func() {
			result := conversion.ConvertFieldType("string")
			Expect(result).To(Equal("string"))
		})

		It("should convert sbyte to byte", func() {
			result := conversion.ConvertFieldType("sbyte")
			Expect(result).To(Equal("byte"))
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

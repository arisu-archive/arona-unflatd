package fbs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
)

var _ = Describe("Schema", func() {
	Context("Field", func() {
		DescribeTable("IsPrimitive",
			func(fieldType string, expected bool) {
				field := fbs.Field{Type: fieldType}
				Expect(field.IsPrimitive()).To(Equal(expected))
			},
			Entry("bool is primitive", "bool", true),
			Entry("int is primitive", "int", true),
			Entry("string is primitive", "string", true),
			Entry("float is primitive", "float", true),
			Entry("custom type is not primitive", "CustomType", false),
			Entry("empty type is not primitive", "", false),
		)
	})

	Context("Schema", func() {
		var schema *fbs.Schema

		BeforeEach(func() {
			schema = &fbs.Schema{
				Namespace: "TestNamespace",
				Tables: []fbs.Table{
					{
						Name: "TestTable",
						Fields: []*fbs.Field{
							{Name: "id", Type: "int"},
							{Name: "name", Type: "string"},
							{Name: "items", Type: "Item", IsArray: true},
						},
					},
				},
				Enums: []fbs.Enum{
					{
						Name:     "TestEnum",
						DataType: "int",
						Values: []fbs.EnumValue{
							{Name: "None", Value: int64(0)},
							{Name: "First", Value: int64(1)},
						},
					},
				},
				RootTypes: []string{"TestTable"},
			}
		})

		It("should have correct structure", func() {
			Expect(schema.Namespace).To(Equal("TestNamespace"))
			Expect(schema.Tables).To(HaveLen(1))
			Expect(schema.Enums).To(HaveLen(1))
			Expect(schema.RootTypes).To(HaveLen(1))
		})

		It("should have correct table structure", func() {
			table := schema.Tables[0]
			Expect(table.Name).To(Equal("TestTable"))
			Expect(table.Fields).To(HaveLen(3))

			// Check field types
			Expect(table.Fields[0].IsPrimitive()).To(BeTrue())
			Expect(table.Fields[1].IsPrimitive()).To(BeTrue())
			Expect(table.Fields[2].IsPrimitive()).To(BeFalse())
			Expect(table.Fields[2].IsArray).To(BeTrue())
		})

		It("should have correct enum structure", func() {
			enum := schema.Enums[0]
			Expect(enum.Name).To(Equal("TestEnum"))
			Expect(enum.DataType).To(Equal("int"))
			Expect(enum.Values).To(HaveLen(2))
			Expect(enum.Values[0].Name).To(Equal("None"))
			Expect(enum.Values[0].Value).To(Equal(int64(0)))
		})
	})
})

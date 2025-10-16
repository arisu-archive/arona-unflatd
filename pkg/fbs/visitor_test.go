package fbs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
)

var _ = Describe("SchemaVisitor", func() {
	var (
		visitor *fbs.SchemaVisitor
		schema  *fbs.Schema
	)

	BeforeEach(func() {
		visitor = fbs.NewSchemaVisitor()
		schema = &fbs.Schema{
			Namespace: "TestNamespace",
			Imports:   []string{"OtherSchema"},
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

	Context("when visiting a schema", func() {
		It("should generate correct FlatBuffer schema text", func() {
			result := visitor.VisitSchema(schema)

			// Check namespace
			Expect(result).To(ContainSubstring("namespace TestNamespace;"))

			// Check imports
			Expect(result).To(ContainSubstring("include \"OtherSchema.fbs\";"))

			// Check table
			Expect(result).To(ContainSubstring("table TestTable {"))
			Expect(result).To(ContainSubstring("id: int;"))
			Expect(result).To(ContainSubstring("name: string;"))
			Expect(result).To(ContainSubstring("[Item];"))

			// Check enum
			Expect(result).To(ContainSubstring("enum TestEnum: int {"))
			Expect(result).To(ContainSubstring("None = 0,"))
			Expect(result).To(ContainSubstring("First = 1"))

			// Check root_type
			Expect(result).To(ContainSubstring("root_type TestTable;"))
		})

		It("should sort imports alphabetically", func() {
			schema.Imports = []string{"Z", "X", "Y"}
			result := visitor.VisitSchema(schema)
			Expect(result).To(ContainSubstring(`include "X.fbs";
include "Y.fbs";
include "Z.fbs";`))
		})
	})

	Context("when visiting a schema with multiple root types", func() {
		BeforeEach(func() {
			schema.RootTypes = append(schema.RootTypes, "AnotherTable")
		})

		It("should generate multiple root_type declarations", func() {
			result := visitor.VisitSchema(schema)
			Expect(result).To(ContainSubstring("root_type TestTable;"))
			Expect(result).To(ContainSubstring("root_type AnotherTable;"))
		})
	})

	Context("when visiting an empty schema", func() {
		BeforeEach(func() {
			schema = &fbs.Schema{
				Namespace: "EmptyNamespace",
			}
		})

		It("should generate minimal valid schema", func() {
			result := visitor.VisitSchema(schema)
			Expect(result).To(ContainSubstring("namespace EmptyNamespace;"))
			Expect(result).NotTo(ContainSubstring("include"))
			Expect(result).NotTo(ContainSubstring("table"))
			Expect(result).NotTo(ContainSubstring("enum"))
			Expect(result).NotTo(ContainSubstring("root_type"))
		})
	})
})

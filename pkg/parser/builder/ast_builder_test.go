package builder_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/builder"
)

var _ = Describe("ASTBuilder", func() {
	var b *builder.ASTBuilder

	BeforeEach(func() {
		b = builder.NewASTBuilder()
	})

	Context("when building a complete AST", func() {
		BeforeEach(func() {
			b.SetNamespace("FlatData")

			// Add enum
			b.AddEnum(&ast.EnumInfo{
				Name:     "TestEnum",
				BaseType: "int",
				Keys:     []string{"None", "First"},
				Values:   []string{"0", "1"},
			})

			// Add struct
			b.BeginStruct("TestStruct")
			b.AddField(&ast.FieldInfo{
				Name:      "Id",
				Type:      "int",
				Modifiers: []string{"public"},
			})
			b.AddMethod(&ast.MethodInfo{
				Name:       "FinishTestStructBuffer",
				Modifiers:  []string{"public"},
				ReturnType: "void",
			})
		})

		It("should build correct FileInfo", func() {
			result := b.Build()

			Expect(result.Namespace).To(Equal("FlatData"))

			// Check enum
			Expect(result.Enums).To(HaveKey("TestEnum"))
			enum := result.Enums["TestEnum"]
			Expect(enum.Keys).To(Equal([]string{"None", "First"}))
			Expect(enum.Values).To(Equal([]string{"0", "1"}))

			// Check struct
			Expect(result.Structs).To(HaveKey("TestStruct"))
			str := result.Structs["TestStruct"]
			Expect(str.Fields).To(HaveLen(1))
			Expect(str.Fields[0].Name).To(Equal("Id"))
			Expect(str.Methods).To(HaveKey("FinishTestStructBuffer"))
		})
	})

	Context("when merging duplicate entries", func() {
		It("should merge enum values correctly", func() {
			enum1 := &ast.EnumInfo{
				Name:   "TestEnum",
				Keys:   []string{"None"},
				Values: []string{"0"},
			}
			enum2 := &ast.EnumInfo{
				Name:   "TestEnum",
				Keys:   []string{"First"},
				Values: []string{"1"},
			}

			b.AddEnum(enum1)
			b.AddEnum(enum2)

			result := b.Build()
			Expect(result.Enums["TestEnum"].Keys).To(Equal([]string{"None", "First"}))
			Expect(result.Enums["TestEnum"].Values).To(Equal([]string{"0", "1"}))
		})
	})
})

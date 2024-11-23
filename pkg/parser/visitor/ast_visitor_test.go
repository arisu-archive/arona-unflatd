package visitor_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/visitor"
)

var _ = Describe("ASTVisitor", func() {
	var v *visitor.ASTVisitor

	BeforeEach(func() {
		v = visitor.NewASTVisitor()
	})

	Context("when visiting AST nodes", func() {
		It("should build correct FileInfo", func() {
			// Visit namespace
			v.VisitNamespace("FlatData")

			// Visit enum
			v.VisitEnum(&ast.EnumInfo{
				Name:     "TestEnum",
				BaseType: "int",
				Keys:     []string{"None", "First"},
				Values:   []string{"0", "1"},
			})

			// Visit struct and its members
			v.VisitStruct(&ast.StructInfo{Name: "TestStruct"})
			v.VisitField("TestStruct", &ast.FieldInfo{
				Name:      "Id",
				Type:      "int",
				Modifiers: []string{"public"},
			})
			v.VisitMethod("TestStruct", &ast.MethodInfo{
				Name:       "FinishTestStructBuffer",
				ReturnType: "void",
			})

			// Get result
			result := v.Result()

			Expect(result.Namespace).To(Equal("FlatData"))
			Expect(result.Enums).To(HaveKey("TestEnum"))
			Expect(result.Structs).To(HaveKey("TestStruct"))
			Expect(result.Structs["TestStruct"].Fields).To(HaveLen(1))
			Expect(result.Structs["TestStruct"].Methods).To(HaveKey("FinishTestStructBuffer"))
		})
	})
})

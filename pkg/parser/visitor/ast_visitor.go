package visitor

import (
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/builder"
)

type ASTVisitor struct {
	builder *builder.ASTBuilder
}

func NewASTVisitor() *ASTVisitor {
	return &ASTVisitor{
		builder: builder.NewASTBuilder(),
	}
}

func (v *ASTVisitor) VisitNamespace(namespace string) {
	v.builder.SetNamespace(namespace)
}

func (v *ASTVisitor) VisitEnum(enumInfo *ast.EnumInfo) {
	v.builder.AddEnum(enumInfo)
}

func (v *ASTVisitor) VisitStruct(structInfo *ast.StructInfo) {
	v.builder.BeginStruct(structInfo.Name)
}

func (v *ASTVisitor) VisitField(structName string, fieldInfo *ast.FieldInfo) {
	v.builder.BeginStruct(structName)
	v.builder.AddField(fieldInfo)
}

func (v *ASTVisitor) VisitMethod(structName string, methodInfo *ast.MethodInfo) {
	v.builder.BeginStruct(structName)
	v.builder.AddMethod(methodInfo)
}

func (v *ASTVisitor) Result() *ast.FileInfo {
	info := v.builder.Build()
	// Reset the builder
	v.builder.Reset()
	return info
}

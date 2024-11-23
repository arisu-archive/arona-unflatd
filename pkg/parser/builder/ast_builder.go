package builder

import (
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
)

type ASTBuilder struct {
	fileInfo      *ast.FileInfo
	currentStruct string
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		fileInfo: &ast.FileInfo{
			Enums:   make(map[string]*ast.EnumInfo),
			Structs: make(map[string]*ast.StructInfo),
		},
	}
}

func (b *ASTBuilder) SetNamespace(namespace string) {
	if b.fileInfo.Namespace == "" {
		b.fileInfo.Namespace = namespace
	}
}

func (b *ASTBuilder) AddEnum(enumInfo *ast.EnumInfo) {
	if _, exists := b.fileInfo.Enums[enumInfo.Name]; !exists {
		b.fileInfo.Enums[enumInfo.Name] = enumInfo
	} else {
		b.fileInfo.Enums[enumInfo.Name].Merge(enumInfo)
	}
}

func (b *ASTBuilder) BeginStruct(name string) {
	b.currentStruct = name
	if _, exists := b.fileInfo.Structs[name]; !exists {
		b.fileInfo.Structs[name] = &ast.StructInfo{
			Name:    name,
			Fields:  make([]*ast.FieldInfo, 0),
			Methods: make(map[string]*ast.MethodInfo),
		}
	}
}

func (b *ASTBuilder) AddField(field *ast.FieldInfo) {
	if structInfo, exists := b.fileInfo.Structs[b.currentStruct]; exists {
		structInfo.Fields = append(structInfo.Fields, field)
	}
}

func (b *ASTBuilder) AddMethod(method *ast.MethodInfo) {
	if structInfo, ok := b.fileInfo.Structs[b.currentStruct]; ok {
		if _, exists := structInfo.Methods[method.Name]; !exists {
			structInfo.Methods[method.Name] = method
		} else {
			structInfo.Methods[method.Name].Merge(method)
		}
	}
}

func (b *ASTBuilder) Build() *ast.FileInfo {
	return b.fileInfo
}

func (b *ASTBuilder) Reset() {
	b.fileInfo = &ast.FileInfo{
		Enums:   make(map[string]*ast.EnumInfo),
		Structs: make(map[string]*ast.StructInfo),
	}
}

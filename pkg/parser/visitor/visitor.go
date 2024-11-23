package visitor

import "github.com/arisu-archive/arona-unflatd/pkg/parser/ast"

// Visitor interface defines methods for visiting different AST nodes.
type Visitor interface {
	VisitNamespace(namespace string)
	VisitEnum(enumInfo *ast.EnumInfo)
	VisitStruct(structInfo *ast.StructInfo)
	VisitField(structName string, fieldInfo *ast.FieldInfo)
	VisitMethod(structName string, methodInfo *ast.MethodInfo)
}

// Element interface represents visitable elements.
type Element interface {
	Accept(v Visitor)
}

package state

import (
	"github.com/arisu-archive/arona-unflatd/pkg/parser/analyzer"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/visitor"
)

// NodeState interface defines how different states should handle node visits.
type NodeState interface {
	Visit(v visitor.Visitor, state *analyzer.ParsingState)
}

// Create concrete states for each node type.
type (
	NamespaceState struct{}
	EnumState      struct{}
	StructState    struct{}
	FieldState     struct{}
	MethodState    struct{}
	EmptyState     struct{}
)

func (*NamespaceState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitNamespace(state.Namespace)
}

func (*EnumState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitEnum(&ast.EnumInfo{
		Name:     state.Name,
		BaseType: state.Enum.BaseType,
		Keys:     state.Enum.Keys,
		Values:   state.Enum.Values,
	})
}

func (*StructState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitStruct(&ast.StructInfo{
		Name:     state.Name,
		BaseList: state.StructBaseList,
	})
}

func (*FieldState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitField(state.Name, &ast.FieldInfo{
		Name:      state.Field.Name,
		Type:      state.Field.Type,
		Modifiers: state.Modifiers,
	})
}

func (*MethodState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitMethod(state.Name, &ast.MethodInfo{
		Name:           state.Method.Name,
		Modifiers:      state.Modifiers,
		ReturnType:     state.Method.ReturnType,
		ParameterNames: state.Method.ParameterNames,
		ParameterTypes: state.Method.ParameterTypes,
	})
}

func (*EmptyState) Visit(_ visitor.Visitor, _ *analyzer.ParsingState) {}

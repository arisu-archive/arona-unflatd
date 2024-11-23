package state

import (
	"github.com/arisu-archive/arona-unflatd/pkg/parser/analyzer"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/visitor"
)

// NodeState interface defines how different states should handle node visits.
type NodeState interface {
	Visit(visitor visitor.Visitor, state *analyzer.ParsingState)
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

func (s *NamespaceState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	if state.Namespace != "" {
		v.VisitNamespace(state.Namespace)
	}
}

func (s *EnumState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitEnum(&ast.EnumInfo{
		Name:     state.Name,
		BaseType: state.Enum.BaseType,
		Keys:     state.Enum.Keys,
		Values:   state.Enum.Values,
	})
}

func (s *StructState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitStruct(&ast.StructInfo{Name: state.Name})
}

func (s *FieldState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitField(state.Name, &ast.FieldInfo{
		Name:      state.Field.Name,
		Type:      state.Field.Type,
		Modifiers: state.Modifiers,
	})
}

func (s *MethodState) Visit(v visitor.Visitor, state *analyzer.ParsingState) {
	v.VisitMethod(state.Name, &ast.MethodInfo{
		Name:           state.Method.Name,
		Modifiers:      state.Modifiers,
		ReturnType:     state.Method.ReturnType,
		ParameterNames: state.Method.ParameterNames,
		ParameterTypes: state.Method.ParameterTypes,
	})
}

func (s *EmptyState) Visit(_ visitor.Visitor, _ *analyzer.ParsingState) {
	// Do nothing for empty state
}

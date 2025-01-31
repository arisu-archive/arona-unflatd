package analyzer

import "github.com/arisu-archive/arona-unflatd/pkg/parser/ast"

type ParsingState struct {
	Namespace          string
	Name               string
	Modifiers          []string
	Field              ast.FieldInfo
	Enum               ast.EnumInfo
	Method             ast.MethodInfo
	PreviousMethodName string
	StructBaseList     string
}

func (state *ParsingState) Reset() {
	state.Namespace = ""
	state.Name = ""
	state.Modifiers = []string{}
	state.Field = ast.FieldInfo{}
	state.Enum = ast.EnumInfo{
		BaseType: "int",
	}
	state.Method = ast.MethodInfo{}
	state.StructBaseList = ""
}

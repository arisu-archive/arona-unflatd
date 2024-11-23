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
}

func (ps *ParsingState) Reset() {
	ps.Namespace = ""
	ps.Name = ""
	ps.Modifiers = []string{}
	ps.Field = ast.FieldInfo{}
	ps.Enum = ast.EnumInfo{
		BaseType: "int",
	}
	ps.Method = ast.MethodInfo{}
}

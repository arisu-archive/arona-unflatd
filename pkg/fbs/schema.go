package fbs

type Schema struct {
	Namespace string
	Imports   []string
	Tables    []Table
	Enums     []Enum
	RootTypes []string
}

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name    string
	Type    string
	IsArray bool
}

type Enum struct {
	Name     string
	DataType string
	Values   []EnumValue
}

type EnumValue struct {
	Name  string
	Value any
}

func (f Field) IsPrimitive() bool {
	primitiveTypes := map[string]bool{
		"bool":    true,
		"byte":    true,
		"sbyte":   true,
		"char":    true,
		"decimal": true,
		"double":  true,
		"float":   true,
		"int":     true,
		"uint":    true,
		"long":    true,
		"ulong":   true,
		"short":   true,
		"ushort":  true,
		"string":  true,
	}
	return primitiveTypes[f.Type]
}
